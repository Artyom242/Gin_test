package handler

import (
	"gin_test_prjct/api/handler/books"
	"gin_test_prjct/api/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

//-

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &Handler{
		DB: db,
	}
	BookHandler := books.NewBookHandler(h.DB)

	r.LoadHTMLGlob("./../templates/*")

	routes := r.Group("/books")
	routes.Use(middleware.CheckAuth)
	routes.POST("/", BookHandler.AddBook)
	routes.GET("/:id", h.GetBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
	routes.GET("/", h.GetBooks)

	r.POST("/login", h.Login)
	r.POST("/reg", h.RegisterUser)
	r.GET("/logout", h.Logout)
}
