package users

import (
	"gin_test_prjct/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h books.handler) Logout(c *gin.Context) {
	var user models.User

	cookie, err := c.Cookie("auth")
	if err != nil {
		log.Println("Error cookie:", err)
	}

	if result := h.DB.Find(&user).Where("token=?", cookie); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
	} else {
		user.AuthToken = ""
	}

	h.DB.Save(&user)

	c.SetCookie(
		"auth",
		"",
		0,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"logout": "success"})
}
