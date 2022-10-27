package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/stephennwac007/HNG-Tsk/controllers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	server := gin.Default()
	server.SetTrustedProxies(nil)

	server.GET("/", controllers.Get)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println("lamba:=", os.Getenv("PORT"))

	server.Run(":" + port)
}
