package api

import (
	"task/api/handler"
	"task/config"
	"task/storage"
	_ "task/api/docs"

	"github.com/gin-gonic/gin"
	 ginSwagger "github.com/swaggo/gin-swagger"
	 swaggerFiles "github.com/swaggo/files"
)


// @description This is a api of the application
func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI) {
	handler := handler.NewHandler(cfg, store)

	url := ginSwagger.URL("swagger/doc.json") 
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.POST("/register", handler.Create)
	r.POST("/auth", handler.Login)
	r.Use(handler.AuthMiddleware())
	{
	r.GET("/user/:name", handler.GetByName)
	r.POST("/phone", handler.CreatePhone)
	r.GET("/phone", handler.GetByPhone)
	r.PUT("/phone", handler.UpdatePhone)
	r.DELETE("/phone/:phone_id", handler.DeletePhone)
	}
}