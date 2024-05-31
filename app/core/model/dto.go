package model

// ==== Message ====

type MessageCreateReq struct {
	Type         EventType `json:"type"`
	TargetId     string    `json:"target_id"`
	Content      string    `json:"content"`
	Quote        string    `json:"quote"`
	Nonce        string    `json:"nonce"`
	TempTargetId string    `json:"temp_target_id"`
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

// ==== Guild ====
type GuildListResp struct {
	WebResult
	Data *GuildResp `json:"data"`
}

type GuildResp struct {
	Guilds []*GuildInfo `json:"items"`
	Meta   PageMeta     `json:"meta"`
}

type GuildInfo struct {
	GuildId          string     `json:"id"`
	Name             string     `json:"name"`
	Topic            string     `json:"topic"`
	UserId           string     `json:"user_id"`
	Icon             string     `json:"icon"`
	NotifyType       NotifyType `json:"notify_type"`
	Region           string     `json:"region"`
	EnableOpen       bool       `json:"enable_open"`
	OpenId           string     `json:"open_id"`
	DefaultChannelId string     `json:"default_channel_id"`
	WelcomeChannelId string     `json:"welcome_channel_id"`
	BoostNum         int        `json:"boost_num"`
	Level            int        `json:"level"`
}

// ==== Channel ====
type ChannelListResp struct {
	WebResult
	Data *ChannelResp `json:"data"`
}

type ChannelResp struct {
	Channels []*ChannelInfo `json:"items"`
	Meta     PageMeta       `json:"meta"`
}

type ChannelInfo struct {
	ChannelId   string      `json:"id"`
	UserId      string      `json:"user_id"`
	ParentId    string      `json:"parent_id"`
	Name        string      `json:"name"`
	Type        ChannelType `json:"type"`
	Level       int         `json:"level"`
	LimitAmount int         `json:"limit_amount"`
	IsCategory  bool        `json:"is_category"`
}
