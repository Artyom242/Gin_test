package users

import (
	"fmt"
	"gin_test_prjct/api/middleware"
	"gin_test_prjct/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type RegForm struct {
	Name           string `json:"name" `
	Surname        string `json:"surname" `
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}

func RegisterUser(c *gin.Context, h *gorm.DB) {
	var dataForm RegForm

	if err := c.ShouldBind(&dataForm); err != nil {
		log.Println("form empty")
		c.JSON(http.StatusBadRequest, &dataForm)
		return
	}

	if dataForm.Name == "" || dataForm.Surname == "" || dataForm.Password == "" || dataForm.RepeatPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Все поля обязательны для заполнения"})
		return
	}

	var user models.User

	hashPass, err := bcrypt.GenerateFromPassword([]byte(dataForm.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Ошибка хэширования:", err)
		return
	}

	tokenString, err := middleware.GenerateToken()
	if err != nil {
		fmt.Println("Ошибка генерации токена:", err)
		return
	}

	user.Name = dataForm.Name
	user.Surname = dataForm.Surname
	user.Password = string(hashPass)
	user.AuthToken = tokenString

	if result := h.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании пользователя"})
		return
	}

	c.SetCookie(
		"auth",
		tokenString,
		3600*24,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, &user)
}
