package handlers

import (
	"TaskReminder/models"
	"TaskReminder/pkg/auth"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecievedUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func LoginAPI(ctx *gin.Context) {
	var receivedUser RecievedUser
	if err := ctx.ShouldBindJSON(&receivedUser); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"error": err,
			})
		return
	}
	var u models.User
	u.SetUserName(receivedUser.Username)
	u.SetPassword(receivedUser.Password)
	if err := u.GetOne("username", receivedUser.Username); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Username not exist",
		})
		return
	}

	if err := auth.ValidatePassword(u.GetPassword(), receivedUser.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Wrong password",
		})
		return
	}
	claim := &auth.Claims{
		ID: u.GetUserID(),
	}
	tokenString, err := auth.GenerateTokenString(*claim)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to generate token string",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"username": receivedUser.Username,
	})
	fmt.Println(`done`)
}

func SignUp(ctx *gin.Context) {
	var signUpUser SignUpUser
	if err := ctx.ShouldBindJSON(&signUpUser); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"error": err,
			})
		return
	}

	var user models.User
	hashedPassword, err := auth.HashPassword(signUpUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"error": "Fail to HashPassword",
			})
		return
	}

	err = user.Create(signUpUser.Username, hashedPassword, signUpUser.Email)
	if err != nil {
		log.Println(err.Error())
		var message string
		if err.Error() == "username is existed" {
			message = err.Error()
		} else {
			message = "Fail to create user"
		}
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"error": message,
			})
		return
	}

	// Hide password
	signUpUser.Password = "*"

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "Sign Up successfully",
			"user":   signUpUser,
		})
}
