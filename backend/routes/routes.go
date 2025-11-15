package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	//inisiasi router 
	router := gin.Default()
	//route register
	router.POST("/api/register", controllers.Register)
	//route login
	router.POST("/api/login", controllers.Login)
	//router users
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)
	//router user create 
	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)
	//router user by id
	router.POST("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUserById)
	return router
}