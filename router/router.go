package router

import (
	"httpserver-temp/handler/cameraalarm"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {

	r := gin.Default()

	route := r.Group("/camera")
	{
		route.POST("/alarm", cameraalarm.CameraAlarm)
	}

	return r
}
