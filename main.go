package main

import (
	"automatic-trade/app"
	"fmt"
)

func main() {
	marketInformationResponse, err := app.GetMarketInformation()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*marketInformationResponse)
}
