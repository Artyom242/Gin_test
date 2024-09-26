package users

import (
	"fmt"
	"gin_test_prjct/api/middleware"
	"gin_test_prjct/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginForm struct {
	Name     string `json:"name" `
	Password string `json:"password"`
}

func (h books.handler) Login(c *gin.Context) {
	var dataForm LoginForm

	if err := c.ShouldBind(&dataForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Плохо сформирован запрос"})
		return
	}

	if dataForm.Name == "" || dataForm.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Все поля обязательны для заполнения"})
		return
	}

	var user models.User

	if result := h.DB.Where("name=?", dataForm.Name).First(&user); result.Error != nil {
		c.JSON(http.StatusNotFound, result.Error)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dataForm.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный пароль"})
		return
	}

	tokenString, err := middleware.GenerateToken()
	if err != nil {
		fmt.Println("Ошибка генерации токена:", err)
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
