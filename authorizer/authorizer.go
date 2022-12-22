package authorizer

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorizeRequestForRoles(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		userRoles := claims["roles"].([]interface{})
		auth := false
		for _, r := range userRoles {
			if contains(roles, r.(string)) {
				auth = true
			}
		}

		if !auth {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		}
	}
}

func ExtractUsername(c *gin.Context) string {
	claims := jwt.ExtractClaims(c)
	username := claims["username"].(string)
	return username
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
