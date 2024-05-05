package action

// ==== Message ====

type MessageCreateReq struct {
	Type         int    `json:"type"`
	TargetId     string `json:"target_id"`
	Content      string `json:"content"`
	Quote        string `json:"quote"`
	Nonce        string `json:"nonce"`
	TempTargetId string `json:"temp_target_id"`
}

type MessageUpdateReq struct {
	MsgId        string `json:"msg_id"`
	Content      string `json:"content"`
	Quote        string `json:"quote"`
	TempTargetId string `json:"temp_target_id"`
}

type MessageDeleteReq struct {
	MsgId string `json:"msg_id"`
}
