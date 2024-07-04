package action

import (
	"encoding/json"

	"github.com/aimerny/kook-go/app/common"
	"github.com/aimerny/kook-go/app/core/helper"
	"github.com/aimerny/kook-go/app/core/model"
	"github.com/sirupsen/logrus"
)

func MessageList() {

}

func MessageDetail() {

}

func MessageSend(req *model.MessageCreateReq) *model.MessageCreateResp {
	response, err := helper.Post(common.BaseUrl+common.V3Url+common.MessageCreate, &req)
	if err != nil {
		return nil
	}
	var result *model.KookResponse[*model.MessageCreateResp]
	err = json.Unmarshal(response, &result)
	if result.Code != 0 {
		return nil
	}
	return result.Data
}

func MessageUpdate(req *model.MessageUpdateReq) error {
	logrus.Println("开始更新消息msg_id:", req.MsgId)
	response, err := helper.Post(common.BaseUrl+common.V3Url+common.MessageUpdate, &req)
	if err != nil {
		return err
	}
	var result *model.KookResponse[interface{}]
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}
	return nil
}

func MessageDelete() {

}

func MessageReactionList() {

}

func MessageAddReaction() {

}

func MessageDeleteReaction() {

}
