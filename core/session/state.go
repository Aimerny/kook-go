package session

/**                                                _________________
 *       获取gateWay     连接ws          收到hello   |    心跳超时     |
 *             |           |                |      |      |         |
 *             v           v                v      |      V         |
 *      INIT  --> GATEWAY -->  WS_CONNECTED --> CONNECTED --> RETRY |
 *       ^        |   ^             |                  ^_______|    |
 *       |        |   |_____________|__________________________|    |
 *       |        |                 |                          |    |
 *       |________|_________________|__________________________|____|
 *
 */

const (
	StatusStart       = "start"
	StatusInit        = "init"
	StatusGateway     = "gateway"
	StatusWSConnected = "wsConnected"
	StatusConnected   = "connected"
	StatusRetry       = "retry"
)

const (
	EventEnterPrefix           = "enter_"
	EventStart                 = "fsmStart"
	EventGotGateway            = "getGateWay"
	EventWsConnected           = "wsConnect"
	EventWsConnectFail         = "wsConnectFail"
	EventHelloReceived         = "helloReceived"
	EventHelloFail             = "helloFail"
	EventHelloGatewayErrFail   = "helloGatewayErrFail"
	EventPongReceived          = "pongReceived"
	EventHeartbeatTimeout      = "heartbeatTimeout"
	EventRetryHeartbeatTimeout = "retryHeartbeatTimeout"
	EventResumeReceivedOk      = "ResumeReceived"
)

const (
	InfiniteRetry = 0
	NoRetry       = -1
)
