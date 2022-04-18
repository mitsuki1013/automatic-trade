package app

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type RequestConfig struct {
	Path      string
	Method    string
	TimeStamp string
}

func NewRequestConfig(path string, method string) *RequestConfig {
	return &RequestConfig{
		Path:      path,
		Method:    method,
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
	}
}

func (c RequestConfig) newAccessSign() string {
	secret := os.Getenv("API_SECRET")

	text := c.TimeStamp + c.Method + c.Path
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func (c RequestConfig) NewPrivateRequest(body io.Reader) (*http.Request, error) {
	uri := os.Getenv("URI")

	req, err := http.NewRequest(c.Method, uri+c.Path, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c RequestConfig) PrivateRequest(req *http.Request) (*http.Response, error) {
	key := os.Getenv("API_KEY")

	req.Header.Set("content-type", "application/json; charset=UTF-8")
	req.Header.Set("ACCESS-KEY", key)
	req.Header.Set("ACCESS-TIMESTAMP", c.TimeStamp)
	req.Header.Set("ACCESS-SIGN", c.newAccessSign())

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
