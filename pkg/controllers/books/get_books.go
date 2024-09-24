package books

import (
	"gin_test_prjct/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetBooks(c *gin.Context) {
	var user models.User
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	//User
	dataCookie, err := c.Cookie("auth")
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	if result := h.DB.Where("token=?", dataCookie).Find(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":  &user,
		"books": &books,
	})
}
