package handler

import (
	"gin_test_prjct/api/handler/books"
	"gin_test_prjct/api/handler/users"
	"gin_test_prjct/api/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, h *gorm.DB) {
	r.LoadHTMLGlob("./../web/templates/*")

	routes := r.Group("/books")
	routes.Use(middleware.CheckAuth)
	routes.POST("/", func(c *gin.Context) { books.AddBook(c, h) })
	routes.GET("/", func(c *gin.Context) { books.GetBooks(c, h) })
	routes.GET("/:id", func(c *gin.Context) { books.GetBook(c, h) })
	routes.PUT("/:id", func(c *gin.Context) { books.UpdateBook(c, h) })
	routes.DELETE("/:id", func(c *gin.Context) { books.DeleteBook(c, h) })

	r.POST("/login", func(c *gin.Context) { users.Login(c, h) })
	r.POST("/reg", func(c *gin.Context) { users.RegisterUser(c, h) })
	r.GET("/logout", func(c *gin.Context) { users.Logout(c, h) })
}
