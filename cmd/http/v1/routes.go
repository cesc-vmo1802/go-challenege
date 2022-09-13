package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-challenege/app_context"
	"go-challenege/docs"
	"go-challenege/features/application/transports/gin_app"
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
	}
}

func privateRoute(sc app_context.AppCtx) func(e *gin.RouterGroup) {
	return func(e *gin.RouterGroup) {

		//TODO: need to add some middleware here to protect your api
		app := e.Group("/applications")
		{
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
