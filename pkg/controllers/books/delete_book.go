package books

import (
	"gin_test_prjct/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	c.Status(http.StatusOK)
}
