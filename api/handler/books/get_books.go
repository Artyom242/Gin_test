package books

import (
	"gin_test_prjct/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetBooks(c *gin.Context, h *gorm.DB) {
	var books []models.Book

	if result := h.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &books)
}
