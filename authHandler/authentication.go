package authHandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(c *gin.Context) {

	token := c.Request.Header["Token"]
	if token == nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Token not present in request",
		})
	}
	fmt.Println(token)

}
