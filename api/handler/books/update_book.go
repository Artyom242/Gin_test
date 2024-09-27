package books

import (
	"gin_test_prjct/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UpdateBookForm struct {
	Title       *string `json:"title"`
	Author      *string `json:"author"`
	Description *string `json:"description"`
}

func UpdateBook(c *gin.Context, h *gorm.DB) {
	id := c.Param("id")
	body := UpdateBookForm{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.First(&book, id); result.Error != nil {
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

	h.Save(&book)

	c.JSON(http.StatusOK, &book)
}
