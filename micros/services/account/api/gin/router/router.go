package router

import (
	"project/micros/services/account/api/gin/handler"

	"github.com/gin-gonic/gin"
)



func NewGinHandler() *gin.Engine {
	router := gin.Default()

	// use cors...
	router.Use()

	// web http handler
	router.POST("registry",handler.RegistryAccount)
	router.POST("registry/tokne",handler.MergeService)



	return router
}