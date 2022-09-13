package gin_app

import (
	"github.com/gin-gonic/gin"
	"go-challenege/app_context"
	"go-challenege/common"
	"go-challenege/features/application/dto"
	"go-challenege/features/application/storage"
	"go-challenege/features/application/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// UpdateApplication
// @Summary 	Update Application
// @Description Update Application
// @Tags 		Applications
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param 		id			path		int		true		"id"
// @Param 		role		body		dto.UpdateApplicationRequest 	true "Update Application"
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/applications/{id} 		[put]
func UpdateApplication(sc app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var form dto.UpdateApplicationRequest
		if err := c.ShouldBind(&form); err != nil {
			panic(err)
		}

		db := sc.GetMgoDB()
		store := storage.NewMongoApplicationStorage(db)
		uc := usecase.NewUpdateApplicationUseCase(store)

		if err := uc.UpdateApplication(c.Request.Context(), objID, &form); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
