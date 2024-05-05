package session

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aimerny/kook-sdk/common"
	"github.com/aimerny/kook-sdk/core/helper"
	"github.com/aimerny/kook-sdk/core/model"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"io"
	"net/url"
	"sync"
	"time"
)

func (s *State) GetGateway(compress bool) error {
	gateway := common.BaseUrl + common.V3Url + common.GateWayUrl

	response, err := helper.Get(gateway)
	if err != nil {
		log.Errorf("Get Gateway faild ! err:%e", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	gatewayInfo := &model.GatewayResp{}
	err = json.Unmarshal(body, gatewayInfo)
	if err != nil {
		return err
	}
	// todo compress use conf
	s.parseUrlWithCompress(gatewayInfo.GatewayInfo.Url, compress)
	s.getGatewaySuccess(s.Conn.Url)
	return nil
}

// 处理url,启动ws连接
func (s *State) parseUrlWithCompress(rawUrl string, compress bool) {
	res, err := url.Parse(rawUrl)
	if err != nil {
		log.Panicf("error url received: %s", rawUrl)
	}
	path := fmt.Sprintf("%s://%s%s", res.Scheme, res.Host, res.Path)
	tokenValue := res.Query().Get("token")
	compressValue := 0
	if compress {
		compressValue = 1
	}
	resUrl := fmt.Sprintf("%s?compress=%d&token=%s", path, compressValue, tokenValue)
	log.Infof("parsed token: %s, result url: %s", tokenValue, resUrl)
	// 设置wsConn
	wsConn := KookWsConn{
		WebConn:     nil,
		Url:         resUrl,
		Token:       tokenValue,
		WsWriteLock: sync.Mutex{},
	}
	s.Conn = &wsConn
}

func (s *State) WsConnect() error {
	wsConn := s.Conn
	wsUrl := wsConn.Url
	if s.SessionId != "" {
		wsUrl = fmt.Sprintf("%s&sn=%d&session_id=%s&resume=1", wsUrl, s.MaxSN, s.SessionId)
	}
	log.Infof("connect to gateway: %s", wsUrl)

	conn, resp, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	log.Infof("connect to gateway resp: %+v", resp)
	if err != nil {
		log.Panicf("connect to ws server failed, %e", err)
	}
	// set ws conn
	s.Conn.WebConn = conn
	s.wsConnectSuccess()
	go func() {
		defer func() {
			conn.Close()
		}()
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Errorf("read msg error: %e", err)
				break
			}
			log.WithField("msg", msg).Trace("websocket data receive")
			s.ReceiveData(msg)
		}
	}()
	return nil
}

func (s *State) sendHeartBeat() error {
	pingPkg := model.NewPing(s.MaxSN)
	if s.Conn != nil {
		data, err := sonic.Marshal(pingPkg)
		if err != nil {
			log.WithError(err).Errorf("marshal ping pkg failed")
			return err
		}

		s.LastPingAt = time.Now()
		log.WithField("ping", string(data)).Info("send Ping")
		err = s.Conn.SendData(data)
		if err != nil {
			return err
		} else {
			s.PongTimeoutChan <- s.LastPingAt.Add(time.Duration(s.Timeout) * time.Second)
		}
	}
	return nil
}

func (s *State) retryHeartBeat() error {
	log.Infof("retry heart beat...")
	err := s.sendHeartBeat()
	if err != nil {
		return err
	}
	return nil
}

func (s *State) StartCheckHeartbeat() error {

	err := s.HeartBeatCron.Every(30).Seconds().Do(s.sendHeartBeat)
	if err != nil {
		log.Errorf("send ping err: %e", err)
		return err
	}
	s.HeartBeatCron.Start()
	log.Infof("Heart beat checker inited")
	return nil
}

// ==== State Change ====

func (s *State) getGatewaySuccess(gateWay string) {
	log.WithField("gateway", gateWay).Info("GetGatewayOk")
	// 流转状态
	s.FSM.Event(context.Background(), EventGotGateway)
}

func (s *State) wsConnectSuccess() {
	log.Info("wsConnectOk")
	s.FSM.Event(context.Background(), EventWsConnected)
}

func (s *State) wsConnectFail() error {
	log.Errorf("ws connect fail")
	s.FSM.Event(context.Background(), EventWsConnectFail)
	return nil
}
