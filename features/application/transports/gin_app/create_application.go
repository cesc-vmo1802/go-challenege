package gin_app

import (
	"github.com/gin-gonic/gin"
	"go-challenege/app_context"
	"go-challenege/common"
	"go-challenege/features/application/dto"
	"go-challenege/features/application/storage"
	"go-challenege/features/application/usecase"
	"net/http"
)

// CreateApplication
// @Summary 	Create Application
// @Description Create Application
// @Tags 		Applications
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Success 	200
// @Param 		Permission		body		dto.CreateApplicationRequest 	true "Create Application"
// @Failure 	400 			{object} 	common.AppError
// @Failure 	404 			{object} 	common.AppError
// @Router 		/api/v1/applications 			[post]
func CreateApplication(sc app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.CreateApplicationRequest

		if err := c.ShouldBind(&form); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := sc.GetMgoDB()
		store := storage.NewMongoApplicationStorage(db)
		uc := usecase.NewCreateApplicationUseCase(store)

		if err := uc.CreateApplication(c.Request.Context(), &form); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
