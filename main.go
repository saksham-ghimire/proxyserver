package main

import (
	"fmt"
	"proxy/authHandler"

	"github.com/gin-gonic/gin"
)

// General Request struct
type Request struct {
	Passphrase string `json:"passphrase"`
	Service    string `json:"service"`
}

func byPassed(c *gin.Context) {
	c.JSON(200, gin.H{
		"meesage": "successfully logged in",
	})

}
func main() {
	authHandler.Load_config()
	fmt.Println(" == Server Started == ")
	r := gin.Default()
	r.Use(authHandler.AuthMiddleWare)
	r.POST("/", byPassed)
	r.Run(":8080")

}
