package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-challenege/app_context"
	"go-challenege/common"
	"go-challenege/features/auth/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if strings.ToLower(parts[0]) != strings.ToLower("Bearer") || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequiredAuth(sc app_context.AppCtx) func(c *gin.Context) {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := sc.GetMgoDB()
		store := storage.NewMongoUserStorage(db)
		aTok := sc.GetATProvider()

		payload, err := aTok.Validate(token)
		if err != nil {
			panic(err)
		}

		objID, err := primitive.ObjectIDFromHex(payload.UserId)
		if err != nil {
			panic(err)
		}
		user, err := store.FindOneByID(c.Request.Context(), objID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		if user.IsBlocked() {
			panic(common.ErrNoPermission(errors.New("status is zero")))
		}

		c.Set(common.RequesterKey, user)
		c.Next()
	}
}
