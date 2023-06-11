package main

import (
	"context"
	"fmt"
	"log"
	"task/api"
	"task/config"
	"task/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	store, err := postgres.NewConnectPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("Ошибка с подключением к postgresql: ", err)
		return
	}
	defer store.CloseDB()
	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	api.NewApi(r, &cfg, store)

	fmt.Println("Server running on port", cfg.ServerHost+cfg.ServerPort)
	err = r.Run(cfg.ServerHost + cfg.ServerPort)
	if err != nil {
		log.Panic("Error listening server: ", err)
		return
	}
}