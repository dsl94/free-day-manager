package routes

import (
	"freeDayManager/authorizer"
	"freeDayManager/controller"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine, userController *controller.UserController, freeDayController *controller.FreeDayController, authMiddleware *jwt.GinJWTMiddleware) {
	app := router.Group("/api")
	app.GET("/refresh_token", authMiddleware.RefreshHandler)
	app.Use(authMiddleware.MiddlewareFunc(), authorizer.AuthorizeRequestForRoles([]string{"ROLE_ADMIN"}))
	{
		app.GET("/admin/users", userController.FindAll)
		app.GET("/admin/users/:id", userController.FindOne)
		app.PUT("/admin/users/:id", userController.Update)
		app.POST("/admin/users", userController.Create)
		app.DELETE("/admin/users/:id", userController.Delete)
		app.GET("/admin/free-days", freeDayController.FindAll)
	}
	app.Use(authMiddleware.MiddlewareFunc(), authorizer.AuthorizeRequestForRoles([]string{"ROLE_ADMIN", "ROLE_USER"}))
	{
		app.POST("/free-days", freeDayController.Create)
		app.GET("/free-days", freeDayController.FindForUser)
	}
}
