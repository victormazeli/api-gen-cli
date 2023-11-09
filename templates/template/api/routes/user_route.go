package routes

import (
	"github.com/gin-gonic/gin"
	"goApiStartetProject/api/handlers"
	"goApiStartetProject/internal/config"
	"gorm.io/gorm"
)

func UserRoute(cfg *config.Config, db *gorm.DB, r *gin.RouterGroup) {
	userHandler := handlers.UserHandler{
		Cfg: cfg,
		Db:  db,
	}

	r.GET("/:id", userHandler.GetUser)
}
