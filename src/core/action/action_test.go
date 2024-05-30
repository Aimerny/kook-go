package action

import (
	"github.com/aimerny/kook-go/common"
	"github.com/aimerny/kook-go/core/helper"
	log "github.com/sirupsen/logrus"
	"testing"
)

func Init() {
	config := common.ReadConfig()
	helper.InitHelper(config.BotToken)
	common.InitLogger()
}

func TestGuildList(t *testing.T) {
	Init()
	log.Infof("%v", GuildList())
}

func TestChannelList(t *testing.T) {
	Init()
	id := GuildList().Data.Guilds[0].GuildId
	list := ChannelList(id)
	log.Info(list)
	log.Info(list.Data.Channels)
}
