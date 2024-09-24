package books

import (
	"fmt"
	"gin_test_prjct/pkg/common/models"
	"gin_test_prjct/pkg/common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type LoginForm struct {
	Name     string `json:"name" `
	Password string `json:"password"`
}

func (h handler) Login(c *gin.Context) {
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

	claims := models.User{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "book-test",
		},
	}

	tokenString, err := utils.GenerateToken(claims)
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
	c.Redirect(http.StatusOK, "/")
}
