package router

import (
	"avito-practice/api/handlers"
	"avito-practice/internal/app"

	"github.com/gorilla/mux"
)

// инициализирует роутер и эндпоинты
//
// принимает переменную приложения для передачи хэндлеров
func SetupRouter(app *app.Application) *mux.Router {
	r := mux.NewRouter()
	handlers := handlers.NewHandler(app)

	// эндпоинт проверки работы сервера
	r.HandleFunc("/ping", handlers.PingHandler).Methods("GET")

	r.HandleFunc("/meetings", handlers.GetProfMeetings).Methods("GET")
	r.HandleFunc("/meetings/filter", handlers.GetProfsStudentMeetings).Methods("GET").
		Queries("student", "{student:[0-9]+}")
	r.HandleFunc("/meeting/{meeting_id:[0-9]+}", handlers.GetMeeting).Methods("GET")
	r.HandleFunc("/meeting/add", handlers.AddMeeting).Methods("POST")

	return r
}
