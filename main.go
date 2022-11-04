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
	server.Use(CORSMiddleware())
	server.SetTrustedProxies(nil)

	server.GET("/", controllers.Get)
	server.POST("/calculate", controllers.MathHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println("lamba:=", os.Getenv("PORT"))

	server.Run(":" + port)
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

			if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(204)
					return
			}

			c.Next()
	}
}