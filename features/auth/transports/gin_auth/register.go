package gin_auth

import (
	"github.com/gin-gonic/gin"
	"go-challenege/app_context"
	"go-challenege/common"
	"go-challenege/features/auth/dto"
	"go-challenege/features/auth/storage"
	"go-challenege/features/auth/usecase"
	"go-challenege/pkg/hash/md5"
	"net/http"
)

// Register
// @Summary 	Register
// @Description Register
// @Tags 		Auth
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Success 	200
// @Param 		Permission		body		dto.CreateUserRequest 	true "Create User"
// @Failure 	400 			{object} 	common.AppError
// @Failure 	404 			{object} 	common.AppError
// @Router 		/api/v1/auth/register 			[post]
func Register(sc app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.CreateUserRequest

		if err := c.ShouldBind(&form); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := sc.GetMgoDB()
		md5Hash := md5.NewMD5Hash()

		store := storage.NewMongoUserStorage(db)
		uc := usecase.NewRegisterUserUseCase(store, md5Hash)

		if err := uc.RegisterUser(c.Request.Context(), &form); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
