package event

import (
	"github.com/aimerny/kook-sdk/core/model"
)

type BaseEventHandler struct{}

func (h *BaseEventHandler) Before(event *model.Event)      {}
func (h *BaseEventHandler) After(event *model.Event)       {}
func (h *BaseEventHandler) DoText(event *model.Event)      {}
func (h *BaseEventHandler) DoPicture(event *model.Event)   {}
func (h *BaseEventHandler) DoVideo(event *model.Event)     {}
func (h *BaseEventHandler) DoFile(event *model.Event)      {}
func (h *BaseEventHandler) DoVoice(event *model.Event)     {}
func (h *BaseEventHandler) DoKMarkDown(event *model.Event) {}
func (h *BaseEventHandler) DoCard(event *model.Event)      {}
func (h *BaseEventHandler) DoSystemMsg(event *model.Event) {}
func (h *BaseEventHandler) DoAll(event *model.Event) {
	h.DoText(event)
	h.DoPicture(event)
	h.DoVideo(event)
	h.DoFile(event)
	h.DoVoice(event)
	h.DoKMarkDown(event)
	h.DoCard(event)
	h.DoSystemMsg(event)
}

type TextEventHandler struct{ BaseEventHandler }
type PictureEventHandler struct{ BaseEventHandler }
type VideoEventHandler struct{ BaseEventHandler }
type FileEventHandler struct{ BaseEventHandler }
type VoiceEventHandler struct{ BaseEventHandler }
type KMarkdownEventHandler struct{ BaseEventHandler }
type CardEventHandler struct{ BaseEventHandler }
type SystemMsgEventHandler struct{ BaseEventHandler }
type AllEventHandler struct{ BaseEventHandler }

func (h *TextEventHandler) DoText(event *model.Event)           {}
func (h *PictureEventHandler) DoPicture(event *model.Event)     {}
func (h *VideoEventHandler) DoVideo(event *model.Event)         {}
func (h *FileEventHandler) DoFile(event *model.Event)           {}
func (h *VoiceEventHandler) DoVoice(event *model.Event)         {}
func (h *KMarkdownEventHandler) DoKMarkDown(event *model.Event) {}
func (h *CardEventHandler) DoCard(event *model.Event)           {}
func (h *SystemMsgEventHandler) DoSystemMsg(event *model.Event) {}
func (h *AllEventHandler) DoAll(event *model.Event)             {}
