package routes

import (
	controller "github.com/Ayush10/PortfoAI/internal/controllers"
	middleware "github.com/Ayush10/PortfoAI/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
    incomingRoutes.Use(middleware.Authenticate())
    incomingRoutes.GET("/users", controller.GetUsers())
    incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
