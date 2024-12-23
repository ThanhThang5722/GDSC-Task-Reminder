package handlers

import (
	"TaskReminder/models"
	"TaskReminder/pkg/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIDFromToken(ctx *gin.Context) (int, error) {
	token, err := ctx.Cookie("jwt_token")
	if token == `` {
		fmt.Println(`null token`)
	}
	if err != nil {
		return 0, err
	}
	claim, err := auth.ParseToken(token)

	if err != nil {
		return 0, err
	}
	return claim.ID, nil
}

func CreatTask(ctx *gin.Context) {
	userID, err := GetIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task models.ReceivedTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = task.InsertTask(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Send a success response
	ctx.JSON(http.StatusOK, gin.H{"message": "Task added successfully!", "task": task})
}
func RemoveTask(ctx *gin.Context) {
	userID, err := GetIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task models.ReceivedTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = task.DeleteTask(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Send a success response
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed successfully!", "task": task})
}

func UpdatePriority(ctx *gin.Context) {
	userID, err := GetIDFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task models.ReceivedTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.UpdateTaskPriority(userID)

	// Send a success response
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated priority successfully!", "task": task})
}
