package books

import (
	"gin_test_prjct/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetBook(c *gin.Context, h *gorm.DB) {
	id := c.Param("id")

	var book models.Book

	if result := h.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &book)
}
