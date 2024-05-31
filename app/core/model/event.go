package model

import "github.com/mitchellh/mapstructure"

type Event struct {
	ChannelType  ChannelEventType `json:"channel_type"`
	EventType    EventType        `json:"type"`
	TargetId     string           `json:"target_id"`
	AuthorId     string           `json:"author_id"`
	Content      string           `json:"content"`
	MsgId        string           `json:"msg_id"`
	MsgTimestamp int              `json:"msg_timestamp"`
	Nonce        string           `json:"nonce"`
	Extra        any              `json:"extra"`
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
	EventType    EventType `json:"type" mapstructure:"event_type"`
	GuildId      string    `json:"guild_id" mapstructure:"guild_id"`
	ChannelName  string    `json:"channel_name" mapstructure:"channel_name"`
	Mention      []string  `json:"mention" mapstructure:"mention"`
	MentionAll   bool      `json:"mention_all" mapstructure:"mention_all"`
	MentionRoles []any     `json:"mention_roles" mapstructure:"mention_roles"`
	MentionHere  bool      `json:"mention_here" mapstructure:"mention_here"`
	Author       User      `json:"author" mapstructure:"author"`
}

type SystemExtra struct {
	SystemEventType string         `json:"type"`
	Body            map[string]any `json:"body"`
}

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
	}
	return nil
}
