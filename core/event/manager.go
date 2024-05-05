package event

import (
	"github.com/aimerny/kook-sdk/core/model"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"time"
)

type EventManager struct {
	queue    *EventQueue
	nextSn   int
	handlers []TypedEventHandler
	started  bool
}
type TypedEventHandler interface {
	Before(event *model.Event)
	After(event *model.Event)
	DoText(event *model.Event)
	DoPicture(event *model.Event)
	DoVideo(event *model.Event)
	DoFile(event *model.Event)
	DoVoice(event *model.Event)
	DoKMarkDown(event *model.Event)
	DoCard(event *model.Event)
	DoSystemMsg(event *model.Event)
	DoAll(event *model.Event)
}

func NewEventManager(queue *EventQueue, nextSn int) *EventManager {
	return &EventManager{
		queue:    queue,
		nextSn:   nextSn,
		handlers: make([]TypedEventHandler, 0),
	}
}

func (manager *EventManager) RegisterHandler(handler []TypedEventHandler) {
	manager.handlers = append(manager.handlers, handler...)
}

func (manager *EventManager) Start() {
	if manager.started {
		return
	}
	manager.started = true
	for {
		if manager.queue.IsEmpty() || manager.queue.MinSn > manager.nextSn {
			log.Trace("empty event list or nextSn < newly sn, sleeping")
			time.Sleep(1 * time.Second)
			continue
		}
		signal := manager.queue.Pop()
		manager.nextSn++
		evt := &model.Event{}
		marshal, err := jsoniter.Marshal(&signal.Data)
		if err != nil {
			log.Errorf("event format failed! data:%v", signal.Data)
		}
		log.Debugf("event:%s", marshal)
		err = jsoniter.Unmarshal(marshal, evt)
		if err != nil {
			log.Errorf("event unmarshal failed! %e", err)
			continue
		}
		go func() {
			for _, h := range manager.handlers {
				h.Before(evt)
				switch evt.EventType {
				case model.EventTypeText:
					h.DoText(evt)
				case model.EventTypePicture:
					h.DoPicture(evt)
				case model.EventTypeVideo:
					h.DoVideo(evt)
				case model.EventTypeFile:
					h.DoFile(evt)
				case model.EventTypeVoice:
					h.DoVoice(evt)
				case model.EventTypeKMarkdown:
					h.DoKMarkDown(evt)
				case model.EventTypeCard:
					h.DoCard(evt)
				case model.EventTypeSystem:
					h.DoSystemMsg(evt)
				}
				h.After(evt)
			}
		}()
	}
}
