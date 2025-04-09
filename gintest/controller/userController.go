package controller

import (
	. "ginprac/model"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	context.HTML(200, "index.html", nil)
}

func Login(context *gin.Context) {
	context.HTML(200, "login.html", nil)
}

func Register(context *gin.Context) {
	switch context.Request.Method {
	case "GET":
		context.HTML(200, "register.html", nil)
	case "POST":
		ua := UserAccount{}
		err := context.ShouldBind(&ua)
		if err != nil {
			context.String(200, err.Error())
			return
		}
		context.JSON(200, gin.H{
			"user":  ua.GetName(),
			"pwd":   ua.GetPwd(),
			"email": ua.Email,
		})
	}
}

func Auth(context *gin.Context) {
	user := context.PostForm("username")
	pwd := context.PostForm("password")
	context.JSON(200, gin.H{
		"user": user,
		"pwd":  pwd,
	})
}

func GetBook(context *gin.Context) {
	context.JSON(200, gin.H{
		"user":  "auser",
		"age":   123,
		"books": []string{"西游记", "三国演义", "国色天香"},
	})
}
