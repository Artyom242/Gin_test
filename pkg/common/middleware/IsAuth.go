package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAuth(c *gin.Context) {
	token, err := c.Request.Cookie("auth")
	if err != nil || token == nil {
		c.JSON(http.StatusOK, gin.H{
			"auth": "token not hounded",
		})
		//c.HTML(http.StatusOK, "login.html", nil)
		c.Abort()
		return
	} else {
		c.Next()
		return
	}
}
