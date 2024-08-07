package action

import (
	"fmt"
	"github.com/aimerny/kook-go/app/common"
	"github.com/aimerny/kook-go/app/core/model"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

func ChannelList(guildId string) *model.ChannelListResp {
	return channelList(guildId, 0, 0)
}

func channelList(guildId string, page, pageSize int) *model.ChannelListResp {
	url := fmt.Sprintf("%s?guild_id=%s", common.ChannelList, guildId)
	if pageSize != 0 {
		url = fmt.Sprintf("%s&page=%d&pageSize=%d", url, page, pageSize)
	}
	resp, err := doGet(url)
	if err != nil {
		log.WithError(err).Error("get channels failed")
	}
	channelListResp := &model.ChannelListResp{}
	err = jsoniter.Unmarshal(resp, &channelListResp)
	if err != nil {
		log.WithError(err).Error("unmarshal resp failed")
	}
	return channelListResp
}

func PageChannelList(guildId string, page, pageSize int) *model.ChannelListResp {
	return channelList(guildId, page, pageSize)
}
