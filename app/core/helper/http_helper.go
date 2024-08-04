package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

var globalToken string = ""

func InitHelper(token string) {
	globalToken = token
}

func Get(url string) (*http.Response, error) {

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Http request construct failed! err: %e", err)
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bot %s", globalToken))
	return client.Do(request)
}

func Post(url string, body interface{}) ([]byte, error) {
	bodyData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(bodyData)

	request, err := http.NewRequest("POST", url, buffer)
	request.Header.Add("Authorization", fmt.Sprintf("Bot %s", globalToken))
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	return postAction(request)
}

func PostWithHeaders(url string, body []byte, headers map[string]string) ([]byte, error) {
	buffer := bytes.NewBuffer(body)
	request, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bot %s", globalToken))
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	return postAction(request)
}

func postAction(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("HTTP POST status code not success,:%d", resp.StatusCode))
	}
	temp, err := io.ReadAll(resp.Body)
	return temp, nil
}

// PostWithJsonBody 发送Post请求到指定url,将body进行json序列化作为请求体
func PostWithJsonBody(url string, body any) {

	client := &http.Client{}
	bodyData, err := jsoniter.Marshal(body)
	if err != nil {
		log.Errorf("marshal body to json failed! body:%v err:%e", body, err)
		return
	}
	log.Tracef("http request body:%s", string(bodyData))
	buffer := bytes.NewBuffer(bodyData)

	request, err := http.NewRequest("POST", url, buffer)
	request.Header.Add("Authorization", fmt.Sprintf("Bot %s", globalToken))
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	// 请求
	resp, err := client.Do(request)
	if err != nil {
		log.Errorf("HTTP POST error: %e", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		log.Errorf("HTTP POST status code not success, %s", respBody)
	}
}
