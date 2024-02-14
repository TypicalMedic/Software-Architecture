package main

import (
	"PAPS-lab5/user_service/api/router"
	"PAPS-lab5/user_service/internal/app"
	"PAPS-lab5/user_service/services/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"PAPS-lab5/user_service/services/configuration"

	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
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

	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	dbconn, err := database.CreateDatabaseConnectionMySQL(connString)
	dbb := sqldblogger.OpenDriver(connString, dbconn.Driver(), loggerAdapter /*, using_default_options*/) // db is STILL *sql.DB

	if err != nil {
		log.Fatal(err)
		return
	}
	// реализуем интерфейс через с полученным подключением
	db = database.NewDatabase(dbb)

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
