package routes

import(
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	//inisiasi router 
	router := gin.Default()
	//route register
	router.POST("/api/register", controllers.Register)
	//route login
	router.POST("/api/login", controllers.Login)

	return router
}