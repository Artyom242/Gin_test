package books

import (
	"gin_test_prjct/pkg/common/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.LoadHTMLGlob("./../templates/*")

	routes := r.Group("/books")
	routes.Use(middleware.CheckAuth)
	routes.POST("/", h.AddBook)
	routes.GET("/:id", h.GetBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
	routes.GET("/", h.GetBooks)

	r.POST("/login", h.Login)
	r.POST("/reg", h.RegisterUser)
	//r.POST("/signup", h.Signup)
	//r.GET("/home", h.Home)
	//r.GET("/premium", h.Premium)
	//r.GET("/logout", h.Logout)
}
