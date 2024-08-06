简体中文 | [English](README_ZH.md)
# KookGo

[![Issues](https://img.shields.io/github/issues/Aimerny/kook-go?style=flat-square)](https://github.com/Aimerny/kook-go/issues)
[![Pull Requests](https://img.shields.io/github/issues-pr/Aimerny/kook-go?style=flat-square)](https://github.com/Aimerny/kook-go/pulls)
[![Release](https://img.shields.io/github/v/release/Aimerny/kook-go?include_prereleases&style=flat-square)](https://github.com/Aimerny/kook-go/releases)
[![Github Release Downloads](https://img.shields.io/github/downloads/Aimerny/kook-go/total?label=Github%20Release%20Downloads&style=flat-square)](https://github.com/Aimerny/kook-go/releases)

基于websocket协议的kook机器人开发sdk

## 快速启动

### 导入模块

```shell
go get github.com/aimerny/kook-go
```

### 简单实现

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

// DoKMarkDown 一个简单的Kook机器人实现,接受到新消息时会发送回对应频道/私聊
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

### [CHANGELOG](./app/CHANGELOG.md)
