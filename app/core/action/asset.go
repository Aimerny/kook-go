package action

import (
	"bytes"
	"github.com/aimerny/kook-go/app/common"
	"github.com/aimerny/kook-go/app/core/model"
	jsoniter "github.com/json-iterator/go"
	"mime/multipart"
)

func AssetCreate(filename string, content []byte) (*model.AssetResp, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	file.Write(content)
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Content-Type"] = writer.FormDataContentType()
	resp, err := doPostWithHeaders(common.AssetCreate, body.Bytes(), headers)
	if err != nil {
		return nil, err
	}
	result := &model.KookResponse[*model.AssetResp]{}
	err = jsoniter.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}
