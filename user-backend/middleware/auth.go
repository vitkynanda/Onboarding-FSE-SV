package middleware

import (
	"go-api/helpers"
	"go-api/usecase/jwt_usecase"

	"github.com/gin-gonic/gin"
)

func JWTAuth(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		userId, err := jwtUsecase.ValidateTokenAndGetUserId(authHeader)
		if err != nil {
			resp := helpers.ResponseError("You are unathorized", err, 401)
			c.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}

		c.Set("user_id", userId)
	}
}
