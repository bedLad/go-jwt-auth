package routes

import (
	"github.com/bedLad/go-jwt-auth/controllers"
	"github.com/bedLad/go-jwt-auth/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users:user_id", controllers.GetUser())
}
