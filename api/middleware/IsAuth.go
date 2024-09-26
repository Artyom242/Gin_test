package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckAuth(c *gin.Context) {
	token, err := c.Cookie("auth")
	log.Println(token)
	if err != nil || token == "" {
		c.JSON(http.StatusOK, gin.H{
			"auth": "token not hounded",
		})
		c.Abort()
		return
	} else {
		c.Next()
		return
	}
}
