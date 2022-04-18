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

type PrivateRequest struct {
	Path      string
	Method    string
	TimeStamp string
}

func NewPrivateRequest(path string, method string) *PrivateRequest {
	return &PrivateRequest{
		Path:      path,
		Method:    method,
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
	}
}

func (c PrivateRequest) NewRequest(body io.Reader) (*http.Request, error) {
	uri := os.Getenv("URI")

	req, err := http.NewRequest(c.Method, uri+c.Path, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c PrivateRequest) DoRequest(req *http.Request) (*http.Response, error) {
	key := os.Getenv("API_KEY")

	req.Header.Set("content-type", "application/json; charset=UTF-8")
	req.Header.Set("ACCESS-KEY", key)
	req.Header.Set("ACCESS-TIMESTAMP", c.TimeStamp)
	req.Header.Set("ACCESS-SIGN", newAccessSign(c.TimeStamp, c.Method, c.Path))

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func newAccessSign(timeStamp string, method string, path string) string {
	secret := os.Getenv("API_SECRET")

	text := timeStamp + method + path
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
