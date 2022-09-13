package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-challenege/app_context"
	"go-challenege/docs"
	"go-challenege/features/application/transports/gin_app"
	"go-challenege/features/auth/transports/gin_auth"
	"go-challenege/middlewares"
	"net/http"
)

func publicRoute(sc app_context.AppCtx) func(e *gin.RouterGroup) {
	return func(e *gin.RouterGroup) {
		e.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
			return
		})

		auth := e.Group("auth")
		{
			auth.POST("/register", gin_auth.Register(sc))
			auth.POST("/login", gin_auth.Login(sc))
		}
	}
}

func privateRoute(sc app_context.AppCtx) func(e *gin.RouterGroup) {
	return func(e *gin.RouterGroup) {

		//TODO: need to add some middleware here to protect your api
		app := e.Group("/applications", middlewares.RequiredAuth(sc))
		{
			app.GET("", gin_app.ListApplication(sc))
			app.GET("/:id", gin_app.DetailApplication(sc))
			app.POST("", gin_app.CreateApplication(sc))
			app.PUT("/:id", gin_app.UpdateApplication(sc))
			app.DELETE("/:id", gin_app.DeleteApplication(sc))
		}
	}
}

func swaggerRoute(sc app_context.AppCtx) func(e *gin.RouterGroup) {
	return func(e *gin.RouterGroup) {
		docs.SwaggerInfo.BasePath = ""
		e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func SetupRoute(sc app_context.AppCtx) func(e *gin.Engine) {
	return func(e *gin.Engine) {
		v1 := e.Group("/api/v1")

		publicRoute(sc)(v1)
		privateRoute(sc)(v1)
		swaggerRoute(sc)(v1)
	}
}
