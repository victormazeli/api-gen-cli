package routes

import (
	"github.com/gin-gonic/gin"
	"goApiStartetProject/internal/config"
	"gorm.io/gorm"
)

func SetupRoute(cfg *config.Config, db *gorm.DB, rg *gin.RouterGroup) {
	UserRoute(cfg, db, rg)

}
