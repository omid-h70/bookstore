package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/omid-h70/bookstore/users-api/app"
	"github.com/omid-h70/bookstore/users-api/utils"
	"log"
)

// main
func main() {
	cfg := utils.Config{
		HttpServerAddress: ":8080",
	}

	fmt.Println("Hi i'm up")

	mux := echo.New()
	server := app.NewApp(&cfg, mux)
	if err := server.Run(); err != nil {
		log.Fatalln("Failed to Start Http Server")
	}
}
