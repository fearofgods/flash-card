package initialize

import (
	"flash-card.juslt.click/global"
	"flash-card.juslt.click/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.AppInfo.Env == "development" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	/// Middleware config
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	/// End middleware config

	/// Route config
	health := new(handlers.HealthController)

	r.GET("/health", health.Status)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(handlers.UserController)
			userGroup.GET("/:id", user.GetUserById)
		}
	}
	/// End route config

	return r
}
