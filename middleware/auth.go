package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/irvanherz/goblog/model"
)

// AuthMiddleware middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
				Status: http.StatusUnauthorized,
				Error:  model.NewRequestError("x", "Invalid auth header"),
			})
			return
		}
		authTokens := strings.Split(authHeader, " ")
		if len(authTokens) != 2 || authTokens[0] != "Bearer" || authTokens[1] == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
				Status: http.StatusUnauthorized,
				Error:  model.NewRequestError("x", "Invalid auth header"),
			})
			return
		}
		var authData model.AuthData
		_, err := jwt.ParseWithClaims(authTokens[1], &authData, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("DEFILATIFAH"), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
				Status: http.StatusUnauthorized,
				Error:  model.NewRequestError("x", err.Error()),
			})
			return
		}
		c.Set("AuthData", authData)
		c.Next()
	}
}
