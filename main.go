package main

import (
	"automatic-trade/app"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file\nError: %v", err)
	}

	permissions, err := app.GetPermission()
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
