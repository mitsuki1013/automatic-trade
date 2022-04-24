package main

import (
	"automatic-trade/app/api"
	"automatic-trade/app/usecase"
	"automatic-trade/client"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file\nError: %v", err)
	}

	r := client.NewClient()
	permissionUseCase := usecase.NewPermissionUseCase(api.NewGetPermissionRequest(r))
	marketInformationUseCase := usecase.NewMarketInformationUseCase(api.NewGetMarketInformationRequest(r))

	permissions, err := permissionUseCase.GetPermission()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*permissions)

	marketInformation, err := marketInformationUseCase.GetMarketInformation()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*marketInformation)
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}
