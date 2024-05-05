package main

import (
	"github.com/aimerny/kook-sdk/common"
	"github.com/aimerny/kook-sdk/core/action"
	"github.com/aimerny/kook-sdk/core/event"
	"github.com/aimerny/kook-sdk/core/model"
	"github.com/aimerny/kook-sdk/core/session"
	log "github.com/sirupsen/logrus"
)

func main() {

	common.InitLogger()
	config := common.ReadConfig()

	globalSession, err := session.CreateSession(config.BotToken, config.Compress)
	if err != nil {
		log.Errorf("%s", err)
	}
	globalSession.RegisterEventHandler(&KMarkDownEventHandler{})
	globalSession.Start()
}

type KMarkDownEventHandler struct {
	event.KMarkdownEventHandler
}

func (h *KMarkDownEventHandler) DoKMarkDown(event *model.Event) {
	content := event.Content
	log.Infof("event:%v", event)
	extra := event.GetUserExtra()
	if extra.Author.Bot {
		log.Warnf("Bot message, skip")
		return
	}
	req := &action.MessageCreateReq{
		Type:     9,
		Content:  content,
		TargetId: event.TargetId,
	}
	action.MessageSend(req)
}
