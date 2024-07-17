package session

import (
	"context"
	"github.com/aimerny/kook-go/app/core/helper"
	"github.com/aimerny/kook-go/app/core/model"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"time"
)

func handle(state *State, signalData *model.Signal) {

	switch signalData.SignalType {
	// hello pack
	case model.SIG_HELLO:
		helloPack := &model.HelloPack{}
		err := mapstructure.Decode(signalData.Data, helloPack)
		if err != nil {
			log.WithField("pack", signalData.Data).Errorf("Hello pack parse failed, error:%e", err)
		}
		handleHello(state, helloPack)
	case model.SIG_PONG:
		// send heart beat to skip this ping wait
		log.Infof("pong!")
		state.HeartBeatPongChan <- true
	case model.SIG_RECONNECT:
		state.EventQueue.Clear()
		state.MaxSN = 0
		helper.StoreSession("", 0)
		state.SessionId = ""
		state.FSM.Event(context.Background(), EventWsConnectFail)
	case model.SIG_EVENT:
		// ignore
		if signalData.SerialNumber < state.MaxSN || (signalData.SerialNumber == 0 && state.MaxSN == 25565) {
			return
		}
		state.EventQueue.Push(signalData)
		state.MaxSN = signalData.SerialNumber
		helper.StoreSession(state.SessionId, state.MaxSN)
	case model.SIG_RESUME_ACK:
		log.Infof("resume session ok.")
	}
}

func handleHello(state *State, pack *model.HelloPack) {
	if pack.Code == model.Success {
		// hello接收成功
		helper.StoreSession(pack.SessionId, state.MaxSN)
		state.SessionId = pack.SessionId
		state.FSM.Event(context.Background(), EventHelloReceived)
	} else if pack.Code == model.ExpiredToken {
		// token超时,重新拉gateway
		log.Errorf("connect to server failed ,gateway token expired. retry get gateway")
		state.FSM.Event(context.Background(), EventHelloGatewayErrFail)
	} else if pack.Code == model.ExpiredSession {
		// session过期,清空sessionId获取新的
		log.Errorf("connect to server failed, gatewat session expired, retry before clear session")
		helper.StoreSession("", 0)
		state.SessionId = ""
		state.FSM.Event(context.Background(), EventHelloFail)
	}
}

// listener
func StartSessionListener(s *State) {
	go func() {
		log.Infof("Listening heart beat!")
		for {
			outTime := <-s.PongTimeoutChan
			log.Debugf("pong time out chan received: %s", outTime)
			// waiting for pong chan's signal
			select {
			// if received pong signal, continue
			case <-s.HeartBeatPongChan:
				// check time
				now := time.Now()
				if now.After(outTime) {
					log.WithField("outTime", outTime).WithField("now", now).Warn("this ping has timed out")
					break
				}
				log.Debug("heart beat check successfully!")
				if s.FSM.Is(StatusRetry) {
					// retry success
					log.Debug("retry success! pong received")
					s.FSM.Event(context.Background(), EventPongReceived)
				}
				break
			case <-time.After(time.Until(outTime)):
				log.Errorf("heart beat time out, try reconnect!")
				if s.FSM.Is(StatusConnected) {
					// clear ping scheduler
					s.HeartBeatCron.Clear()
					log.Debugf("heart beat check scheduler stoped")
					go s.FSM.Event(context.Background(), EventHeartbeatTimeout)
				} else if s.FSM.Is(StatusRetry) {
					log.Infof("reconnect timeout, state turn to get gateway.Clear heart beat scheduler")
					go s.FSM.Event(context.Background(), EventRetryHeartbeatTimeout)
				}
				break
			}
		}
	}()
	// received signal handle listener
	go func() {
		log.Infof("recived event listener start.")
		for {
			sign := <-s.ReceiveQueue
			log.Debug("receive signal: %v", sign)
			handle(s, sign)
		}
	}()
}
