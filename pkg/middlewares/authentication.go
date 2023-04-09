package middlewares

import (
	"net/http"

	tokenjwt "github.com/afrizal423/go-gin-secure-newbie/pkg/utils/tokenJWT"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tokenjwt.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", tokenjwt.ExtractToken(c))
		c.Next()
	}
}
