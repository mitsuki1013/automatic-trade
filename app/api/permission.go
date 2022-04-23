package api

import (
	"automatic-trade/client"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Permission string

type GetPermissionRequest struct {
	client *client.Client
}

func NewGetPermissionRequest(client *client.Client) GetPermissionRequest {
	return GetPermissionRequest{client: client}
}

func (g GetPermissionRequest) GetPermission() (*[]Permission, error) {
	req, err := g.client.NewRequest(http.MethodGet, "/v1/me/getpermissions", nil)
	if err != nil {
		return nil, err
	}

	res, err := g.client.DoRequest(req)
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
