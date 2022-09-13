package gin_app

import (
	"github.com/gin-gonic/gin"
	"go-challenege/app_context"
	"go-challenege/common"
	"go-challenege/features/application/storage"
	"go-challenege/features/application/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// DeleteApplication
// @Summary 	Delete Application
// @Description Delete Application
// @Tags 		Applications
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param 		id			path		int		true		"id"
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/applications/{id} 		[delete]
func DeleteApplication(sc app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := sc.GetMgoDB()
		store := storage.NewMongoApplicationStorage(db)
		uc := usecase.NewDeleteApplicationUseCase(store)

		if err = uc.DeleteApplication(c.Request.Context(), objID); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
