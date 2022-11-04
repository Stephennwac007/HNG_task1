package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"slackUsername": "stephen_bukz",
		"backend":       true,
		"age":           24,
		"bio":           "a MERN stack bot🤖 that's trying to be Proficient with Golang, with constant cruise catching",
	})
}

type Question struct {
	OpsType string `json:"operation_type"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
}

type Reponse struct {
	SlackUsername string `json:"slackUsername"`
	Result        int    `json:"result"`
	OpsType       string `json:"operation_type"`
}

func MathHandler(ctx *gin.Context) {
	var question Question
	if err := ctx.Bind(&question); err != nil {
		return
	}
	result, operator := Check(question)

	res := Reponse{
		SlackUsername: "stephen_bukz",
		Result:        result,
		OpsType:       operator,
	}

	ctx.IndentedJSON(http.StatusCreated, res)
}

func Check(data Question) (int, string) {
	var result int
	var operator_type string

	math := map[string][]string{
		"addition":       {"add", "addition", "plus", "sum"},
		"subtraction":    {"subtract", "subtraction", "minus"},
		"multiplication": {"multiply", "multiplication", "times"},
		"division":       {"divide", "division", "over"},
	}

	for _, val := range math["addition"] {
		if strings.Contains(data.OpsType, val) {
			result = data.X + data.Y
			operator_type = "addition"
		}
	}

	for _, val := range math["multiplication"] {
		if strings.Contains(data.OpsType, val) {
			result = data.X * data.Y
			operator_type = "multiplication"
		}
	}

	for _, val := range math["subtraction"] {
		if strings.Contains(data.OpsType, val) {
			result = data.X - data.Y
			operator_type = "subtraction"
		}
	}

	for _, val := range math["division"] {
		if strings.Contains(data.OpsType, val) {
			result = data.X / data.Y
			operator_type = "division"
		}
	}

	
	return result, operator_type
}
