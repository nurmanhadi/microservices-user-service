package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"user-service/pkg/enum"
	"user-service/pkg/response"
	"user-service/pkg/security"
	"user-service/src/config/env"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

func ErrorHandling() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err

			if validationErr, ok := err.(validator.ValidationErrors); ok {
				var values []string
				for _, fieldErr := range validationErr {
					value := fmt.Sprintf("field %s is %s %s", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
					values = append(values, value)
				}
				arguments := strings.Join(values, ", ")
				ctx.JSON(http.StatusBadRequest, map[string]any{
					"error": arguments,
					"path":  ctx.Request.URL.Path,
				})
				return
			} else if responseStatusException, ok := err.(*response.ErrorResponse); ok {
				ctx.JSON(responseStatusException.Code, map[string]any{
					"error": responseStatusException.Message,
					"path":  ctx.Request.URL.Path,
				})
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, map[string]any{
					"error": err.Error(),
					"path":  ctx.Request.URL.Path,
				})
				return
			}
		}
	}
}
func JwtValidation(role []enum.ROLE) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Except(401, "missing or invalid authorization token")
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(env.CONF.JWT.Access.Secret), nil
		})
		if err != nil || !token.Valid {
			response.Except(401, "invalid or expired token")
			return
		}
		claims := token.Claims.(security.JwtClaims)
		for _, x := range role {
			if claims.Role != x {
				response.Except(403, "insufficient role")
				return
			}
		}
		ctx.Next()
	}
}
