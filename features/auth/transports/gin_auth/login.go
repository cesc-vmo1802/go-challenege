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

// Login
// @Summary 	Login
// @Description Login
// @Tags 		Auth
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Success 	200
// @Param 		Permission		body		dto.LoginUserRequest 	true "Login User"
// @Failure 	400 			{object} 	common.AppError
// @Failure 	404 			{object} 	common.AppError
// @Router 		/api/v1/auth/login 			[post]
func Login(sc app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.LoginUserRequest

		if err := c.ShouldBind(&form); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := sc.GetMgoDB()

		aTok := sc.GetATProvider()
		rTok := sc.GetRTProvider()
		hasher := md5.NewMD5Hash()

		store := storage.NewMongoUserStorage(db)
		uc := usecase.NewLoginUserUseCase(store, hasher, aTok, rTok)
		data, err := uc.Login(c.Request.Context(), &form)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
		return
	}
}
