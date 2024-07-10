package action

import (
	"errors"
	jsoniter "github.com/json-iterator/go"

	"github.com/aimerny/kook-go/app/common"
	"github.com/aimerny/kook-go/app/core/model"
)

func MessageList() {

}

func MessageDetail() {

}

func MessageSend(req *model.MessageCreateReq) (*model.MessageCreateResp, error) {
	response, err := doPost(common.MessageCreate, req)
	if err != nil {
		return nil, err
	}
	var result *model.KookResponse[*model.MessageCreateResp]
	err = jsoniter.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, errors.New(result.Message)
	}
	return result.Data, nil
}

func MessageUpdate(req *model.MessageUpdateReq) error {
	response, err := doPost(common.MessageUpdate, req)
	if err != nil {
		return err
	}
	var result *model.KookResponse[interface{}]
	err = jsoniter.Unmarshal(response, &result)
	if err != nil {
		return err
	}
	if result.Code != 0 {
		return errors.New(result.Message)
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
