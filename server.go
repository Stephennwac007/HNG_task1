package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stephennwac007/HNG-Tsk/controllers"
)

func main()  {
	server := gin.Default()

	server.GET("/user", controllers.Get)
	server.Run(":9090")
}