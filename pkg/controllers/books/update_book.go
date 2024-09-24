package books

import (
	"gin_test_prjct/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateBook struct {
	Title       *string `json:"title"`
	Author      *string `json:"author"`
	Description *string `json:"description"`
}

func (h handler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	body := UpdateBook{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	if body.Title != nil {
		book.Title = *body.Title
	}
	if body.Author != nil {
		book.Author = *body.Author
	}
	if body.Description != nil {
		book.Description = *body.Description
	}

	h.DB.Save(&book)

	c.JSON(http.StatusOK, &book)
}
