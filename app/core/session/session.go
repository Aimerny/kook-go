package session

import (
	"bytes"
	"compress/zlib"
	"context"
	"errors"
	"github.com/aimerny/kook-go/app/core/event"
	"github.com/aimerny/kook-go/app/core/helper"
	"github.com/aimerny/kook-go/app/core/model"
	"github.com/avast/retry-go/v4"
	"github.com/gorilla/websocket"
	"github.com/jasonlvhit/gocron"
	"github.com/looplab/fsm"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/signal"
	"sync"
	"time"
)

type StatusParam struct {
	StartTime  int
	MaxTime    int
	FirstDelay int
	Retry      int
	MaxRetry   int
}

type KookSession struct {
	Token string
	State *State
}

type State struct {
	Compress          bool
	Conn              *KookWsConn
	SessionId         string
	FSM               *fsm.FSM // 状态机
	StatusParams      map[string]*StatusParam
	PongTimeoutChan   chan time.Time
	LastPingAt        time.Time
	Timeout           int
	ReceiveQueue      chan *model.Signal
	HeartBeatPongChan chan bool
	HeartBeatCron     *gocron.Scheduler
	MaxSN             int
	EventQueue        *event.EventQueue
	EventManager      *event.EventManager
}

type KookWsConn struct {
	Url         string
	Token       string
	WebConn     *websocket.Conn
	WsWriteLock sync.Mutex
}

// CreateSession 创建Session
func CreateSession(token string, compress bool) (*KookSession, error) {
	// 初始化Helper层
	helper.InitHelper(token)

	session := &KookSession{
		Token: token,
	}
	session.State = CreateState(compress)
	return session, nil
}

func CreateState(compress bool) *State {
	state := &State{}
	state.StatusParams = map[string]*StatusParam{
		StatusInit:        {StartTime: 0, MaxTime: 60, FirstDelay: 1, MaxRetry: InfiniteRetry},
		StatusGateway:     {StartTime: 1, MaxTime: 32, FirstDelay: 2, MaxRetry: 2},
		StatusWSConnected: {StartTime: 6, MaxTime: 0, FirstDelay: 0, MaxRetry: NoRetry},
		StatusConnected:   {StartTime: 30, MaxTime: 30, FirstDelay: 0, MaxRetry: NoRetry},
		StatusRetry:       {StartTime: 0, MaxTime: 8, FirstDelay: 4, MaxRetry: 2},
	}

	state.FSM = fsm.NewFSM(
		StatusStart,
		fsm.Events{
			// 启动时间
			{Name: EventStart, Src: []string{StatusStart}, Dst: StatusInit},
			// 获取网关事件
			{Name: EventGotGateway, Src: []string{StatusInit}, Dst: StatusGateway},
			// 连接ws事件
			{Name: EventWsConnected, Src: []string{StatusGateway}, Dst: StatusWSConnected},
			// 连接ws失败事件
			{Name: EventWsConnectFail, Src: []string{StatusWSConnected, StatusGateway, StatusRetry, StatusInit, StatusConnected}, Dst: StatusInit},
			// 接收到hello包事件
			{Name: EventHelloReceived, Src: []string{StatusWSConnected}, Dst: StatusConnected},
			// 接受hello包失败事件
			{Name: EventHelloFail, Src: []string{StatusWSConnected}, Dst: StatusGateway},
			// 接受到错误的hello包事件
			{Name: EventHelloGatewayErrFail, Src: []string{StatusWSConnected}, Dst: StatusInit},
			// 接受到pong包
			{Name: EventPongReceived, Src: []string{StatusConnected, StatusWSConnected}, Dst: StatusConnected},
			// 心跳超时
			{Name: EventHeartbeatTimeout, Src: []string{StatusConnected}, Dst: StatusRetry},
			// 重试心跳超时
			{Name: EventRetryHeartbeatTimeout, Src: []string{StatusRetry}, Dst: StatusGateway},
			// 恢复消费完成
			{Name: EventResumeReceivedOk, Src: []string{StatusWSConnected, StatusConnected}, Dst: StatusConnected},
		},
		fsm.Callbacks{
			EventEnterPrefix + StatusInit: func(_ context.Context, e *fsm.Event) {
				state.Retry(e, func() error { return state.GetGateway(compress) }, nil)
			},
			EventEnterPrefix + StatusGateway: func(_ context.Context, e *fsm.Event) {
				state.Retry(e, func() error { return state.WsConnect() }, func() error { return state.wsConnectFail() })
			},
			EventEnterPrefix + StatusWSConnected: func(_ context.Context, e *fsm.Event) {

			},
			EventEnterPrefix + StatusConnected: func(_ context.Context, e *fsm.Event) {
				state.StartCheckHeartbeat()
			},
			EventEnterPrefix + StatusRetry: func(_ context.Context, e *fsm.Event) {
				// if turn to retry status,we should send heart beat twice.
				state.Retry(e, func() error { state.retryHeartBeat(); return errors.New("retry next ping action") }, func() error { return state.sendHeartBeat() })
			},
		},
	)
	state.ReceiveQueue = make(chan *model.Signal)
	state.PongTimeoutChan = make(chan time.Time)
	state.HeartBeatPongChan = make(chan bool)
	state.EventQueue = event.NewEventQueue()
	state.SessionId, state.MaxSN = helper.ReloadSession()
	state.HeartBeatCron = gocron.NewScheduler()
	state.EventManager = event.NewEventManager(state.EventQueue, state.MaxSN+1)
	state.Compress = compress
	go state.EventManager.Start()
	// default timeout time is 6 seconds
	state.Timeout = 6

	return state
}

func (s *KookSession) RegisterEventHandler(handler event.TypedEventHandler) {
	s.State.RegisterEventHandler([]event.TypedEventHandler{handler})
}

func (s *KookSession) RegisterEventHandlers(handler []event.TypedEventHandler) {
	s.State.RegisterEventHandler(handler)
}

func (s *State) RegisterEventHandler(handler []event.TypedEventHandler) {
	s.EventManager.RegisterHandler(handler)
}

func (s *KookSession) Start() {
	s.State.Start()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			log.Println("connect interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := s.State.Conn.WebConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write msg close:", err)
				return
			}
			return
		}
	}
}

func (s *State) Start() {
	if s.Conn == nil {
		s.FSM.SetState(StatusInit)
		s.Retry(nil, func() error { return s.GetGateway(s.Compress) }, nil)

	} else {
		s.FSM.SetState(StatusGateway)
		s.Retry(nil, func() error { return s.WsConnect() }, func() error { return s.wsConnectFail() })
	}
	go StartSessionListener(s)
}

func (s *State) Retry(event *fsm.Event, handler func() error, errorHandler func() error) {
	log.Infof("retry by event:%v", event)
	log.Infof("now session state is: %s", s.FSM.Current())

	startTime := s.StatusParams[s.FSM.Current()].StartTime
	maxTime := s.StatusParams[s.FSM.Current()].MaxTime
	firstDelay := s.StatusParams[s.FSM.Current()].FirstDelay
	maxRetry := s.StatusParams[s.FSM.Current()].MaxRetry

	//如果传了参数就用传的
	if event != nil {
		if len(event.Args) > 0 {
			if param, ok := event.Args[0].(*StatusParam); ok {
				if param.StartTime > 0 {
					startTime = param.StartTime
				}
				if param.MaxTime > 0 {
					maxTime = param.MaxTime
				}
				if param.FirstDelay > 0 {
					firstDelay = param.FirstDelay
				}
				if param.MaxRetry != 0 {
					maxRetry = param.MaxRetry
				}

			}
		}
	}
	time.Sleep(time.Second * time.Duration(startTime))

	// 不需要指数回退重试
	if maxRetry == NoRetry {
		err := handler()
		if err != nil {
			log.Errorf("retry failed, handler:%s", helper.GetFunctionName(handler()))
			if errorHandler() != nil {
				errorHandler()
			}
		}
		return
	}

	err := retry.Do(
		handler,
		retry.DelayType(retry.BackOffDelay),
		retry.Delay(time.Second*time.Duration(firstDelay)),
		retry.MaxDelay(time.Second*time.Duration(maxTime)),
		retry.Attempts(uint(maxRetry)),
		retry.OnRetry(func(n uint, err error) { log.WithError(err).Infof("try %d times call function %s", n, handler) }),
	)

	if err != nil && errorHandler != nil {
		errorHandler()
	}

}

func (s *State) ReceiveData(data []byte) (error, []byte) {
	if data == nil {
		return nil, nil
	}
	if s.Compress == true {
		b := bytes.NewReader(data)
		r, err := zlib.NewReader(b)
		if err != nil {
			return err, nil
		}

		data, err = io.ReadAll(r)
		if err != nil {
			log.Error(err)
			return err, nil
		}
	}

	signalData := model.ParseSignal(data)
	if signalData == nil {
		log.Warnf("Signal invalid! data:%v", data)
	} else {
		log.WithField("signalData", signalData).Debug("receive signalData from server")
		s.ReceiveQueue <- signalData
	}
	return nil, data
}

func (c *KookWsConn) SendData(data []byte) error {
	c.WsWriteLock.Lock()
	defer c.WsWriteLock.Unlock()
	return c.WebConn.WriteMessage(websocket.TextMessage, data)
}
