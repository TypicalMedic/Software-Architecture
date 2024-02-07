package main

import (
	"avito-practice/api/router"
	"avito-practice/internal/app"
	"avito-practice/services/database"
	"fmt"
	"log"
	"net/http"

	"avito-practice/services/configuration"
)

func main() {
	// создаем новую конфигурацию из yaml файла
	config, err := configuration.NewConfig("config.yml")
	if err != nil {
		log.Fatal(err)
		return
	}

	// создаем переменную интерфейса
	var db database.IDatabase
	// собираем строку подключения из конфигурации
	connString := fmt.Sprint(config.Database.Username, ":", config.Database.Password,
		"@tcp(", config.Database.Host, ":", config.Database.Port, ")/", config.Database.Database, "?parseTime=true")
	// подключаемся к бд по строке
	dbconn, err := database.CreateDatabaseConnectionMySQL(connString)
	if err != nil {
		log.Fatal(err)
		return
	}
	// реализуем интерфейс через с полученным подключением
	db = database.NewDatabase(dbconn)

	app, err := app.InitApp(db, *config)
	if err != nil {
		log.Fatal(err)
		return
	}
	// создаем роутер с настройкой
	router := router.SetupRouter(&app)

	log.Print("App started!")
	log.Fatal(http.ListenAndServe(fmt.Sprint(config.Server.Host, ":", config.Server.Port), router))
}
