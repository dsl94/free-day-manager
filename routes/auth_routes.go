package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	router.GET("/refresh_token", authMiddleware.RefreshHandler)
}
