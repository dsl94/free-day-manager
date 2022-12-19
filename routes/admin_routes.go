package routes

import (
	"freeDayManager/authorizer"
	"freeDayManager/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine, userController *user.UserController, authMiddleware *jwt.GinJWTMiddleware) {
	admin := router.Group("/api/admin")
	admin.GET("/refresh_token", authMiddleware.RefreshHandler)
	admin.Use(authMiddleware.MiddlewareFunc(), authorizer.AuthorizeRequestForRoles([]string{"ROLE_ADMIN"}))
	{
		admin.GET("/users", userController.FindAll)
	}
}
