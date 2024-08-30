package middleware

import (
	"crowdfunding/auth"
	"crowdfunding/helper"
	"crowdfunding/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(authJwt auth.JwtService, userService service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			authFailedMiddleware(ctx)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authJwt.ValidateToken(tokenString)
		if err != nil {
			authFailedMiddleware(ctx)
			return
		}
		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			authFailedMiddleware(ctx)
		}
		userID := int(payload["user_id"].(float64))
		user, err := userService.FindById(userID)
		if err != nil {
			authFailedMiddleware(ctx)
			return
		}
		ctx.Set("user", user)
	}
}
func authFailedMiddleware(ctx *gin.Context) {
	response := helper.UnAuthorized("Unauthorized")
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
