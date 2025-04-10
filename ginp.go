package main

import (
	. "ginprac/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	// router
	var getmap = map[string]func(context *gin.Context){
		"index":    Index,
		"/":        Index,
		"books":    GetBook,
		"register": Register,
		"login":    Login,
	}
	var postmap = map[string]func(context *gin.Context){
		"auth":     Auth,
		"register": Register,
		"/":        Register,
	}

	for k, v := range getmap {
		router.GET(k, v)
	}

	for k, v := range postmap {
		router.POST(k, v)
	}
	router.Run(":8090")
}
