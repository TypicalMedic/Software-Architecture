package router

import (
	"PAPS-lab5/meeting_service/api/handlers"
	"PAPS-lab5/meeting_service/internal/app"

	"github.com/gorilla/mux"
)

// инициализирует роутер и эндпоинты
//
// принимает переменную приложения для передачи хэндлеров
func SetupRouter(app *app.Application) *mux.Router {
	r := mux.NewRouter()
	handlers := handlers.NewHandler(app)

	r.HandleFunc("/meetings", handlers.GetProfMeetings).Methods("GET").
		Queries("prof", "{prof:[0-9]+}")
	r.HandleFunc("/meetings/filter", handlers.GetProfsStudentMeetings).Methods("GET").
		Queries("student", "{student:[0-9]+}", "prof", "{prof:[0-9]+}")
	r.HandleFunc("/meeting/{meeting_id:[0-9]+}", handlers.GetMeeting).Methods("GET").
		Queries("prof", "{prof:[0-9]+}")
	r.HandleFunc("/meeting/add", handlers.AddMeeting).Methods("POST")

	return r
}
