package model

type WebResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PageMeta struct {
	Page      int `json:"page"`
	PageTotal int `json:"page_total"`
	PageSize  int `json:"page_size"`
	Total     int `json:"total"`
}

// GatewayInfo Get gateway resp
type GatewayInfo struct {
	Url string `json:"url"`
}

type GatewayResp struct {
	WebResult
	GatewayInfo GatewayInfo `json:"data"`
}
