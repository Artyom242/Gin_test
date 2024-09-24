package main

import (
	"gin_test_prjct/pkg/common/config"
	"gin_test_prjct/pkg/common/db"
	"gin_test_prjct/pkg/controllers/books"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	port := ":" + cnf.Port
	dbUrl := cnf.DBUrl

	log.Printf("Подключение к базе данных по адресу: %s", dbUrl)
	log.Printf("Сервер запустится на порту: %s", port)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)

	r.Run(port)
}
