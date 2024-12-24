package handlers

import (
	"TaskReminder/models"
	"TaskReminder/pkg/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderHomePage(ctx *gin.Context) {
	token, err := ctx.Cookie("jwt_token")
	if token == `` {
		fmt.Println(`token khong truy cap duoc\n`)
	}
	fmt.Println(token)
	if err != nil {
		fmt.Println(`Bug in cookie: `, err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Please login First at /login",
		})
		ctx.Redirect(http.StatusMovedPermanently, "/login")
		ctx.Abort()
		return
	}
	claim, err := auth.ParseToken(token)

	if err != nil {
		fmt.Println(`Bug in Parse`)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Please login First /login",
		})
		ctx.Redirect(http.StatusMovedPermanently, "/login")
		ctx.Abort()
		return
	}

	var user models.User
	user.GetOne(`userID`, claim.ID)

	doFirst, err1 := models.GetGroupTask(`DoFirst`, user.GetUserID())
	doLater, err2 := models.GetGroupTask(`DoLater`, user.GetUserID())
	delegate, err3 := models.GetGroupTask(`Delegate`, user.GetUserID())
	eliminate, err4 := models.GetGroupTask(`Eliminate`, user.GetUserID())
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Print(err1)
	}
	ctx.HTML(http.StatusOK, "main/base-layout.html", gin.H{
		"BrainName": "GDSC Task Reminder",
		"UserName":  user.GetUserName(),
		"doFirst":   doFirst.ListTask,
		"doLater":   doLater.ListTask,
		"delegate":  delegate.ListTask,
		"eliminate": eliminate.ListTask,
	})
}
func RenderLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusAccepted, "main/login-layout.html", gin.H{
		"BrainName": "GDSC Task Reminder",
	})
}
func RenderSignUpPage(ctx *gin.Context) {
	ctx.HTML(http.StatusAccepted, "main/signup-layout.html", gin.H{
		"BrainName": "GDSC Task Reminder",
	})
}
