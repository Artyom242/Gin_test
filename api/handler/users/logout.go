package users

import (
	"gin_test_prjct/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Logout(c *gin.Context, h *gorm.DB) {
	var user models.User

	cookie, err := c.Cookie("auth")
	if err != nil {
		log.Println("Error cookie:", err)
	}

	if result := h.Find(&user).Where("token=?", cookie); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
	} else {
		user.AuthToken = ""
	}

	h.Save(&user)

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
