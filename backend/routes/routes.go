package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	//inisiasi router 
	router := gin.Default()

	//set up cors
		router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	//route register
	router.POST("/api/register", controllers.Register)
	//route login
	router.POST("/api/login", controllers.Login)
	//route users
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)
	//route user create 
	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)
	//route user by id
	router.POST("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUserById)
	//route user update
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
	//route delete
	router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)
	return router
}