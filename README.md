[简体中文](README_ZH.md) | English
# KookGo

A kook robot development sdk based on websocket protocol

## QuickStart

### Import module

```shell
go get github.com/aimerny/kook-go
```

### Simple Implement
```go
func main() {

	common.InitLogger()
	globalSession, err := session.CreateSession("你的KookToken", true)
	if err != nil {
		log.Errorf("%s", err)
	}
	globalSession.RegisterEventHandler(&MyEventHandler{})
	globalSession.Start()
}

type MyEventHandler struct {
	event.BaseEventHandler
}

// DoKMarkDown A simple Kook robot implementation that sends new messages back to the corresponding channel/private chat
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
```

## [CHANGELOG](./app/CHANGELOG.md)
