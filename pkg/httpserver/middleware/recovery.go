package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-challenege/common"
	"go-challenege/pkg/i18n"
)

func Recovery(i18n *i18n.AppI18n) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("content-type", "application/json")
				language := c.GetHeader("")

				if ve, ok := err.(validator.ValidationErrors); ok {
					appVE := common.HandleValidationErrors(language, i18n, ve)
					c.AbortWithStatusJSON(appVE.StatusCode, appVE)
					return
				}

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, common.HandleAppError(language, i18n, appErr))
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
