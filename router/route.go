package router

import (
	"goChat/docs"
	"goChat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/Test", service.Test)

	return r
}
