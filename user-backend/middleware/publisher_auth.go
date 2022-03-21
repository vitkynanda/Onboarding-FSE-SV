package middleware

import (
	"go-api/helpers"
	"go-api/usecase/jwt_usecase"

	"github.com/gin-gonic/gin"
)

func PublisherAuth(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		
		userId, role, err := jwtUsecase.ValidateTokenAndGetRole(authHeader)
		if err != nil {
			resp := helpers.ResponseError("You are unathorized", "Invalid token", 401)
			c.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}
		

		if (!(role == "admin" || role == "publihsher")) {
			resp := helpers.ResponseError("Forbidden Access", "You have no access to do this action", 403)
			c.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}

		c.Set("user_id", userId)
		
	}
}
