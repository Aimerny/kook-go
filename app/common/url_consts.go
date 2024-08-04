package common

const (
	BaseUrl    = "https://www.kookapp.cn/api"
	V3Url      = "/v3"
	GateWayUrl = "/gateway/index"

	// message uris

	MessageList           = "/message/list"
	MessageView           = "/message/view"
	MessageCreate         = "/message/create"
	MessageUpdate         = "/message/update"
	MessageDelete         = "/message/delete"
	MessageReactionList   = "/message/reaction-list"
	MessageAddReaction    = "/message/add-reaction"
	MessageDeleteReaction = "/message/delete-reaction"

	// channel uris

	ChannelList       = "/channel/list"
	ChannelView       = "/channel/view"
	ChannelCreate     = "/channel/create"
	ChannelUpdate     = "/channel/update"
	ChannelDelete     = "/channel/delete"
	ChannelUserList   = "/channel/user-list"
	ChannelMoveUser   = "/channel/move-user"
	ChannelRoleIndex  = "/channel-role/index"
	ChannelRoleCreate = "/channel-role/create"
	ChannelRoleUpdate = "/channel-role/update"
	ChannelRoleSync   = "/channel-role/sync"
	ChannelRoleDelete = "/channel-role/delete"

	// guild uris

	GuildList         = "/guild/list"
	GuildView         = "/guild/view"
	GuildUserList     = "/guild/user-list"
	GuildNickname     = "/guild/nickname"
	GuildLeave        = "/guild/leave"
	GuildKickout      = "/guild/kickout"
	GuildMuteList     = "/guild-mute/list"
	GuildMuteCreate   = "/guild-mute/create"
	GuildMuteDelete   = "/guild-mute/delete"
	GuildBoostHistory = "/guild-boost/history"

	// bot offline

	OfflineBot = "/user/offline"

	// Asset about
	AssetCreate = "/asset/create"
)
