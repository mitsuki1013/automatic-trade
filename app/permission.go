package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type permission string

func GetPermission() (*[]permission, error) {
	path := "/v1/me/getpermissions"
	requestConfig := NewRequestConfig(path, http.MethodGet)

	req, err := requestConfig.NewPrivateRequest(nil)
	if err != nil {
		return nil, err
	}

	res, err := requestConfig.PrivateRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var permissions []permission
	if err := json.Unmarshal(body, &permissions); err != nil {
		return nil, err
	}

	return &permissions, nil
}
