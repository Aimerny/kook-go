package action

import (
	"fmt"
	"github.com/aimerny/kook-go/common"
	"github.com/aimerny/kook-go/core/model"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

func GuildList() *model.GuildListResp {
	return guildList(common.GuildList)
}

func guildList(url string) *model.GuildListResp {
	resp, err := doGet(url)
	if err != nil {
		log.WithError(err).Error("get guild failed")
	}
	guildResp := &model.GuildListResp{}
	err = jsoniter.Unmarshal(resp, &guildResp)
	if err != nil {
		log.WithError(err).Error("unmarshal resp failed")
	}
	return guildResp
}

func PageGuildList(page, pageSize int) *model.GuildListResp {
	if pageSize == 0 {
		return GuildList()
	}
	url := fmt.Sprintf("%s?page=%d&page_size=%d", common.GuildList, page, pageSize)
	return guildList(url)
}
