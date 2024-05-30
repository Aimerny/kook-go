package action

import (
	"github.com/aimerny/kook-go/app/common"
	"github.com/aimerny/kook-go/app/core/model"
)

func MessageList() {

}

func MessageDetail() {

}

func MessageSend(req *model.MessageCreateReq) {
	go doPost(common.MessageCreate, req)
}

func MessageUpdate() {

}

func MessageDelete() {

}

func MessageReactionList() {

}

func MessageAddReaction() {

}

func MessageDeleteReaction() {

}
