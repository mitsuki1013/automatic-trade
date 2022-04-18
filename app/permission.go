package app

import (
	"encoding/json"
	"io/ioutil"
)

type Permission string

type GetPermissionRequest struct {
	privateRequest *PrivateRequest
}

func NewGetPermissionRequest(privateRequest *PrivateRequest) GetPermissionRequest {
	return GetPermissionRequest{privateRequest: privateRequest}
}

func (g GetPermissionRequest) GetPermission() (*[]Permission, error) {
	req, err := g.privateRequest.NewRequest(nil)
	if err != nil {
		return nil, err
	}

	res, err := g.privateRequest.DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var permissions []Permission
	if err := json.Unmarshal(body, &permissions); err != nil {
		return nil, err
	}

	return &permissions, nil
}
