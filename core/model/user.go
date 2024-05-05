package model

type User struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Nickname       string `json:"nickname"`
	IdentifyNum    string `json:"identify_num"`
	Online         bool   `json:"online"`
	Bot            bool   `json:"bot"`
	Status         int    `json:"status"`
	Avatar         string `json:"avatar"`
	VipAvatar      string `json:"vip_avatar"`
	MobileVerified bool   `json:"mobile_verified"`
	Roles          []int  `json:"roles"`
}

const StatusBan = 10

type Guild struct {
	Id               string     `json:"id"`
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
	Roles            []Role     `json:"roles"`
	Channels         []Channel  `json:"channels"`
}
type NotifyType int

const (
	NotifyTypeServer    NotifyType = 0
	NotifyTypeAll       NotifyType = 1
	NotifyTypeMentioned NotifyType = 2
	NotifyTypeIgnore    NotifyType = 3
)

type Role struct {
	RoleId      int    `json:"role_id"`
	Name        string `json:"name"`
	Color       int    `json:"color"`
	Position    int    `json:"position"`
	Hoist       int    `json:"hoist"`
	Mentionable int    `json:"mentionable"`
	Permissions int    `json:"permissions"`
}

type Channel struct {
	Id                   string      `json:"id"`
	Name                 string      `json:"name"`
	UserId               string      `json:"user_id"`
	GuildId              string      `json:"guild_id"`
	Topic                string      `json:"topic"`
	IsCategory           bool        `json:"is_category"`
	ParentId             string      `json:"parent_id"`
	Level                int         `json:"level"`
	SlowMode             int         `json:"slow_mode"`
	Type                 ChannelType `json:"type"`
	PermissionOverwrites []any       `json:"permission_overwrites"`
	PermissionUsers      []any       `json:"permission_users"`
	PermissionSync       int         `json:"permission_sync"`
	HasPassword          bool        `json:"has_password"`
}
type ChannelType int

const (
	ChannelTypeText  ChannelType = 0
	ChannelTypeVoice ChannelType = 1
)
