package main

import (
	"freeDayManager/database"
	"freeDayManager/routes"
	"freeDayManager/security"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	db := database.Connect()
	database.Migrate(db)
	defer db.Close()

	roleController := InitRoleController(db)
	userController := InitUserController(db, roleController.RoleService.RoleRepository)
	freeDayController := InitFreeDayController(db, userController.UserService.UserRepository)

	router := gin.Default()
	router.Use(cors.Default()) // Allowed origins *

	var authMiddleware = security.CreateAuthMiddleware(&userController.UserService)
	router.POST("/api/login", authMiddleware.LoginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	routes.AuthRoutes(router, authMiddleware)
	routes.AppRoutes(router, &userController, &freeDayController, authMiddleware)

	err := router.Run()
	if err != nil {
		panic(err)
	}

}
