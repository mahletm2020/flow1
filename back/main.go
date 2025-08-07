package main

import (
	"flow1/config"
	"flow1/routes"
	 "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main (){
	r:= gin.Default()
	r.Use(cors.Default())
	config.CoonectDB()
	routes.SetupRouter(r)
	r.Run(":8080")
}