package action

import (
	"github.com/aimerny/kook-go/common"
	"github.com/aimerny/kook-go/core/helper"
	log "github.com/sirupsen/logrus"
	"io"
)

func doGet(actionUrl string) ([]byte, error) {
	resp, err := helper.Get(common.BaseUrl + common.V3Url + actionUrl)
	if err != nil {
		log.Errorf("action failed! err:%e", err)
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("action read body failed! err:%e", err)
		return nil, err
	}
	return data, nil
}

func doPost(actionUrl string, body any) {
	helper.PostWithJsonBody(common.BaseUrl+common.V3Url+actionUrl, body)
}
