package handlers

import (
	"TaskReminder/models"
	"TaskReminder/pkg/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatTask(ctx *gin.Context) {
	token, err := ctx.Cookie("jwt_token")
	if token == `` {
		fmt.Println(`token khong truy cap duoc\n`)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claim, err := auth.ParseToken(token)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task models.ReceivedTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.InsertTask(claim.ID)

	// Send a success response
	ctx.JSON(http.StatusOK, gin.H{"message": "Task added successfully!", "task": task})
}
