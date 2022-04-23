package api

import (
	"automatic-trade/client"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type MarketInformation struct {
	ProductCode string `json:"product_code"`
	MarketType  string `json:"market_type"`
	Alias       string `json:"alias,omitempty"`
}

type GetMarketInformationRequest struct {
	client *client.Client
}

func NewGetMarketInformationRequest(client *client.Client) GetMarketInformationRequest {
	return GetMarketInformationRequest{client: client}
}

func (g GetMarketInformationRequest) GetMarketInformation() (*[]MarketInformation, error) {
	req, err := g.client.NewRequest(http.MethodGet, "/v1/markets", nil)
	if err != nil {
		return nil, err
	}

	res, err := g.client.DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var marketInformation []MarketInformation
	if err := json.Unmarshal(body, &marketInformation); err != nil {
		return nil, err
	}

	return &marketInformation, nil
}
