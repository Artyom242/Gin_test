package books

import (
	"gin_test_prjct/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler.Handler) GetBooks(c *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &books)
}
