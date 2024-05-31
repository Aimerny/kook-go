package main

import (
	"github.com/aimerny/kook-go/app/common"
	"github.com/aimerny/kook-go/app/core/action"
	"github.com/aimerny/kook-go/app/core/event"
	"github.com/aimerny/kook-go/app/core/model"
	"github.com/aimerny/kook-go/app/core/session"
	log "github.com/sirupsen/logrus"
)

func main() {

	common.InitLogger()
	globalSession, err := session.CreateSession("Your kook-go bot token", true)
	if err != nil {
		log.Errorf("%s", err)
	}
	globalSession.RegisterEventHandler(&MyEventHandler{})
	globalSession.Start()
}

type MyEventHandler struct {
	event.BaseEventHandler
}

// DoKMarkDown A simple example to process kmarkdown. It will repeat content from origin message.
func (h *MyEventHandler) DoKMarkDown(event *model.Event) {
	content := event.Content
	log.Infof("event:%v", event)
	extra := event.GetUserExtra()
	if extra.Author.Bot {
		log.Warnf("Bot message, skip")
		return
	}
	req := &model.MessageCreateReq{
		Type:     9,
		Content:  "Repeat by kook bot:" + content,
		TargetId: event.TargetId,
	}
	action.MessageSend(req)
}
