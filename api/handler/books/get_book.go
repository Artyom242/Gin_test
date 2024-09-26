package books

import (
	"gin_test_prjct/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h books.handler) GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &book)
}
