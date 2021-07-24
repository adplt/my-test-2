package middlewares

import (
	"errors"
	"fmt"

	configs "project/app/configs"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		configs.Block{
			Try: func() {
				tokenString := c.Request.Header.Get("Authorization")
				_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if jwt.GetSigningMethod(jwt.SigningMethodHS256.Name) != token.Method {
						return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
					}

					return []byte(token.Raw), nil
				})

				if err != nil {
					configs.Throw(err)
				}

				c.Next()
			},
			Catch: func(e error) {
				c.AbortWithStatusJSON(500, gin.H{
					"status":  "FAILED",
					"message": e.Error(),
					"data":    nil,
				})
				configs.FancyHandleError(e)
			},
		}.Do()
	}
}
