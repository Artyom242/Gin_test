package main

import (
	"gin_test_prjct/api/handler"
	"gin_test_prjct/internal/config"
	"gin_test_prjct/internal/database"
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
	h := database.Init(dbUrl)

	handler.RegisterRoutes(r, h)

	r.Run(port)
}
