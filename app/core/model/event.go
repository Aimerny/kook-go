package model

import (
	"github.com/mitchellh/mapstructure"
)

type Event struct {
	ChannelType  ChannelEventType `json:"channel_type"`
	EventType    EventType        `json:"type"`
	TargetId     string           `json:"target_id"`
	AuthorId     string           `json:"author_id"`
	Content      string           `json:"content"`
	MsgId        string           `json:"msg_id"`
	MsgTimestamp int              `json:"msg_timestamp"`
	Nonce        string           `json:"nonce"`
	Extra        interface{}      `json:"extra"`
}

type ChannelEventType string

const (
	ChannelEventTypeGroup     ChannelEventType = "GROUP"
	ChannelEventTypePerson    ChannelEventType = "PERSON"
	ChannelEventTypeBroadcast ChannelEventType = "BROADCAST"
)

type EventType int

const (
	EventTypeText      EventType = 1
	EventTypePicture   EventType = 2
	EventTypeVideo     EventType = 3
	EventTypeFile      EventType = 4
	EventTypeVoice     EventType = 8
	EventTypeKMarkdown EventType = 9
	EventTypeCard      EventType = 10
	EventTypeSystem    EventType = 255
)

type UserExtra struct {
	EventType    EventType     `json:"type" mapstructure:"event_type"`
	GuildId      string        `json:"guild_id" mapstructure:"guild_id"`
	ChannelName  string        `json:"channel_name" mapstructure:"channel_name"`
	Mention      []string      `json:"mention" mapstructure:"mention"`
	MentionAll   bool          `json:"mention_all" mapstructure:"mention_all"`
	MentionRoles []interface{} `json:"mention_roles" mapstructure:"mention_roles"`
	MentionHere  bool          `json:"mention_here" mapstructure:"mention_here"`
	Author       User          `json:"author" mapstructure:"author"`
}

type SystemExtra struct {
	SystemEventType string                 `json:"type"`
	Body            map[string]interface{} `json:"body"`
}

// system extra Body
// ==== 用户相关事件 ====
// 用户加入语音频道
type JoinedChannel struct {
	UserId    string `json:"user_id" mapstructure:"user_id"`
	ChannelId string `json:"channel_id" mapstructure:"channel_id"`
	JoinedAt  int    `json:"joined_at" mapstructure:"joined_at"`
}

// 用户退出语音频道
type ExitedChannel struct {
	UserId    string `json:"user_id" mapstructure:"user_id"`
	ChannelId string `json:"channel_id" mapstructure:"channel_id"`
	ExitedAt  int    `json:"exited_at" mapstructure:"exited_at"`
}

// 用户信息更新
type UserUpdated struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username" mapstructure:"username"`
	Avatar   int    `json:"avatar" mapstructure:"avatar"`
}

// 自己新加入服务器
type SelfJoinedGuild struct {
	GuildId string `json:"guild_id" mapstructure:"guild_id"`
	State   string `json:"state" mapstructure:"state"`
}

// 自己退出服务器
type SelfExitedGuild struct {
	GuildId string `json:"guild_id" mapstructure:"guild_id"`
}

// Card 消息中的 Button 点击事件
type MessageBtnClick struct {
	GuildId  string `json:"guild_id" mapstructure:"guild_id"`
	MsgId    string `json:"msg_id" mapstructure:"msg_id"`
	TargetId string `json:"target_id" mapstructure:"target_id"`
	UserId   string `json:"user_id" mapstructure:"user_id"`
	UserInfo User   `json:"user_info" mapstructure:"user_info"`
	Value    string `json:"value" mapstructure:"value"`
}

// ==== 用户相关事件 ==== 结束

func (e *Event) GetUserExtra() *UserExtra {
	if e.EventType != EventTypeSystem {
		extra := &UserExtra{}
		mapstructure.Decode(e.Extra, extra)
		return extra
	}
	return nil
}

func (e *Event) GetSystemExtra() *SystemExtra {
	if e.EventType == EventTypeSystem {
		extra := &SystemExtra{}
		mapstructure.Decode(e.Extra, extra)
		return extra
	}
	return nil
}
