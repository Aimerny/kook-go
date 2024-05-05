package model

type WebResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// === 获取Gateway ===
type GatewayInfo struct {
	Url string `json:"url"`
}

type GatewayResp struct {
	WebResult
	GatewayInfo GatewayInfo `json:"data"`
}
