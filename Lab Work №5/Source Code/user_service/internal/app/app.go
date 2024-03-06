package app

import (
	"PAPS-lab5/user_service/services/configuration"
	"PAPS-lab5/user_service/services/database"
)

// структура приложения, где хранится его основная информация
type Application struct {
	Database      database.IDatabase
	Configuration configuration.Config
}

// конструктор приложения
//
// принимает интерфейс бд, конфигурацию
func InitApp(database database.IDatabase, config configuration.Config) (Application, error) {
	a := Application{
		Database:      database,
		Configuration: config,
	}
	return a, nil
}
