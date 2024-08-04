package action

import (
	"bytes"
	"github.com/aimerny/kook-go/app/common"
	"github.com/aimerny/kook-go/app/core/model"
	jsoniter "github.com/json-iterator/go"
	"io"
	"mime/multipart"
)

func AssetCreate(filename string, content []byte) (*model.AssetResp, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := writer.CreateFormFile("uploadfile", filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(file, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	writer.Close()

	headers := make(map[string]string)
	headers["Content-Type"] = writer.FormDataContentType()
	resp, err := doPostWithHeaders(common.AssetCreate, body, headers)
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
