package routes

import (
	"flow1/controllers"

	"github.com/gin-gonic/gin"
)
func SetupRouter(router *gin.Engine){
	router.POST("/api/register" ,controllers.Register)
	router.POST("/api/login" ,controllers.Login)

}