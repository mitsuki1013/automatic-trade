package main

import (
	"automatic-trade/app"
	"automatic-trade/app/usecase"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file\nError: %v", err)
	}

	r := app.NewPrivateRequest("/v1/me/getpermissions", http.MethodGet)
	permissionUseCase := usecase.NewPermissionUseCase(app.NewGetPermissionRequest(r))

	permissions, err := permissionUseCase.GetPermission()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*permissions)
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}
