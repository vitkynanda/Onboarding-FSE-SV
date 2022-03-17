package middleware

import (
	"fmt"
	"go-api/helpers"
	"go-api/usecase/jwt_usecase"

	"github.com/gin-gonic/gin"
)

func JWTAuth(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println(authHeader)
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader)
		if err != nil {
			resp := helpers.ResponseError("You are unathorized", err, 403)
			c.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}
		c.Set("user_id", userId)
	}
}
