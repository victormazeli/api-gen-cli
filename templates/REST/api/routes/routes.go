package routes

import (
	"github.com/gin-gonic/gin"
	"goApiStartetProject/api/handlers"
	"goApiStartetProject/internal/config"
	"goApiStartetProject/pkg/ApiResponse"
	"gorm.io/gorm"
)

func SetupRoute(env *config.Env, db *gorm.DB, routeGroup *gin.RouterGroup) {
	// Health Route
	routeGroup.GET("/healthz", func(c *gin.Context) {
		ApiResponse.SendSuccess(c, "Health")
	})

	// User related routes
	userHandler := handlers.UserHandler{
		Env: env,
		Db:  db,
	}

	routeGroup.GET("/:id", userHandler.GetUser)

}
