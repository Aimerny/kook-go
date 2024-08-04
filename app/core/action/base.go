package action

import (
	"github.com/aimerny/kook-go/app/common"
	"github.com/aimerny/kook-go/app/core/helper"
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

func doPost(actionUrl string, body any) ([]byte, error) {
	response, err := helper.Post(common.BaseUrl+common.V3Url+actionUrl, &body)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func doPostWithHeaders(actionUrl string, body []byte, headers map[string]string) ([]byte, error) {
	return helper.PostWithHeaders(common.BaseUrl+common.V3Url+actionUrl, body, headers)
}
