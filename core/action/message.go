package action

import (
	"github.com/aimerny/kook-sdk/common"
)

func MessageList() {

}

func MessageDetail() {

}

func MessageSend(req *MessageCreateReq) {
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
