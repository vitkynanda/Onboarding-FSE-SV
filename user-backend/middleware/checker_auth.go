package middleware

import (
	"go-api/helpers"
	"go-api/usecase/jwt_usecase"

	"github.com/gin-gonic/gin"
)

func CheckerAuth(jwtUsecase jwt_usecase.JwtUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		userId, role, err := jwtUsecase.ValidateTokenAndGetRole(authHeader)
		
		if err != nil {
			resp := helpers.ResponseError("You are unathorized", "Invalid token", 401)
			c.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}
	
		if (!(role == "admin" || role == "checker")) {
			resp := helpers.ResponseError("Forbidden Access", "You have no access to do this action", 403)
			c.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}

		productId := c.Param("id")
		errProduct := jwtUsecase.CheckProductData(productId, "checker")

		if (productId != "") && (errProduct != nil) {
			if (errProduct.Error() == "record not found") {
				resp := helpers.ResponseError("Data not found", errProduct.Error(), 404)
				c.AbortWithStatusJSON(resp.StatusCode, resp)
				return
			} else {
				resp := helpers.ResponseError("Forbidden Access", errProduct.Error(), 403)
				c.AbortWithStatusJSON(resp.StatusCode, resp)
				return
			}
		} 

		c.Set("user_id", userId)
		
	}
}
