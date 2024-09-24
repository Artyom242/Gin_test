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

type RegForm struct {
	Name           string `json:"name" `
	Surname        string `json:"surname" `
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}

func (h handler) RegisterUser(c *gin.Context) {
	var dataForm RegForm

	if err := c.ShouldBind(&dataForm); err != nil {
		c.JSON(http.StatusBadRequest, dataForm)
		return
	}

	if dataForm.Name != "" || dataForm.Surname != "" || dataForm.Password != "" || dataForm.RepeatPassword != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Все поля обязательны для заполнения"})
		return
	}

	var user models.User

	hashPass, err := bcrypt.GenerateFromPassword([]byte(dataForm.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Ошибка хэширования:", err)
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

	user.Name = dataForm.Name
	user.Surname = dataForm.Surname
	user.Password = string(hashPass)
	user.Name = dataForm.Name
	user.AuthToken = tokenString

	if result := h.DB.Create(&user); result.Error != nil {
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

	c.Redirect(http.StatusOK, "/")
}
