package httpx

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, handler *userHandler) {
	r.POST("/", handler.CreateUser)
	r.GET("/", handler.FindUsers)
}
