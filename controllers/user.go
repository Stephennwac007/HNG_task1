package controllers

import "github.com/gin-gonic/gin"

func Get(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"slackUsername": "stephen_bukz",
		"backend": true,
		"age": 24,
		"bio": "a MERN stack bot🤖 that's trying to be Proficient with Golang, with constant cruise catching",
	})
}