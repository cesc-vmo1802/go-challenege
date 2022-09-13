package gin_app

import (
	"github.com/gin-gonic/gin"
	"go-challenege/app_context"
	"go-challenege/common"
	"go-challenege/features/application/dto"
	"go-challenege/features/application/storage"
	"go-challenege/features/application/usecase"
	"go-challenege/pkg/paging"
	"net/http"
)

// ListApplication
// @Summary 	List Application
// @Description List Application
// @Tags 		Applications
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param 		offset			query		int		true		"offset"
// @Param 		limit			query		int		false		"limit"
// @Param 		name			query		string 	false		"Application name"
// @Param 		description		query		string 	false		"Application description"
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/applications 		[get]
func ListApplication(sc app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter dto.Filter
		var page paging.Paging

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&page); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		page.Fulfill()

		db := sc.GetMgoDB()
		store := storage.NewMongoApplicationStorage(db)
		uc := usecase.NewListApplicationUseCase(store)

		data, err := uc.List(c.Request.Context(), &page, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, page, filter))
		return
	}
}
