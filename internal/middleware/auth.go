package middleware

import (
	"strings"

	"github.com/BramAristyo/go-pos-mawish/pkg/service_errors"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		auth := c.GetHeader("Authorization")

		token := strings.Split(auth, " ")
		if auth == "" || len(token) < 2 {
			err = service_errors.TokenRequired
		} else {
			// TODO FUCKNG
		}
	}
}
