package routes

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.)
	incomingRoutes.POST("users/login", controllers.)
}