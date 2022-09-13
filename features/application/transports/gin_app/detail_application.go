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

// DetailApplication
// @Summary 	Detail Application
// @Description Detail Application
// @Tags 		Applications
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param 		id			path		int		true		"id"
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/applications/{id} 		[get]
func DetailApplication(sc app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := sc.GetMgoDB()
		store := storage.NewMongoApplicationStorage(db)
		uc := usecase.NewDetailApplicationUseCase(store)

		data, err := uc.DetailApplication(c.Request.Context(), objID)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
		return
	}
}
