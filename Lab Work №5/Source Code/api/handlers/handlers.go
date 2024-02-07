package handlers

import (
	"avito-practice/internal/app"
	"avito-practice/services/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	app *app.Application
}

func NewHandler(a *app.Application) Handler {
	return Handler{app: a}
}

// Всегда возвращает 200 (для проверки работоспособности сервера)
func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetProfsStudentMeetings(w http.ResponseWriter, r *http.Request) {
	var msg string      // ответ сервера
	vars := mux.Vars(r) //получаем переменные запроса

	studId, err := strconv.Atoi(vars["student"])
	if err != nil {
		msg = "Unable to parse id: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	headerContentTtype := r.Header.Get("Content-Type")
	// проверяем соответсвтвие типа содержимого запроса
	if headerContentTtype != "application/json" {
		msg = "Incorrect content type!"
		log.Print(msg)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(msg)
		return
	}
	// декодируем тело запроса
	var reqB BaseRequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&reqB)
	if err != nil {
		msg = "Failed to decode request body: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := GetMeetingsResultBody{
		Meetings: []Meeting{},
	}
	m := h.app.Database.GetProfsStudentMeetings(reqB.ProfessorId, studId)
	for _, mt := range m {
		st, _ := h.app.Database.GetStudent(mt.StudentParticipantId)
		res.Meetings = append(res.Meetings,
			Meeting{
				Id:          mt.Id,
				Name:        mt.Name,
				Description: mt.Description,
				MeetingTime: mt.MeetingTime,
				StudentParticipant: Student{
					Id:           st.Id,
					Name:         st.Name,
					Surname:      st.Surname,
					Middlename:   st.Middlename,
					Cource:       st.Cource,
					ProjectTheme: st.ProjectTheme,
				},
				IsOnline: mt.IsOnline})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Получает список сегментов, в которых находится пользователь по его id
func (h *Handler) GetProfMeetings(w http.ResponseWriter, r *http.Request) {
	var msg string // ответ сервера

	headerContentTtype := r.Header.Get("Content-Type")
	// проверяем соответсвтвие типа содержимого запроса
	if headerContentTtype != "application/json" {
		msg = "Incorrect content type!"
		log.Print(msg)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(msg)
		return
	}
	// декодируем тело запроса
	var reqB BaseRequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqB)
	if err != nil {
		msg = "Failed to decode request body: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := GetMeetingsResultBody{
		Meetings: []Meeting{},
	}
	m := h.app.Database.GetProfMeetings(reqB.ProfessorId)
	for _, mt := range m {
		st, _ := h.app.Database.GetStudent(mt.StudentParticipantId)
		res.Meetings = append(res.Meetings,
			Meeting{
				Id:          mt.Id,
				Name:        mt.Name,
				Description: mt.Description,
				MeetingTime: mt.MeetingTime,
				StudentParticipant: Student{
					Id:           st.Id,
					Name:         st.Name,
					Surname:      st.Surname,
					Middlename:   st.Middlename,
					Cource:       st.Cource,
					ProjectTheme: st.ProjectTheme,
				},
				IsOnline: mt.IsOnline})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Добавляет пользователя в бд
func (h *Handler) AddMeeting(w http.ResponseWriter, r *http.Request) {

	var msg string // ответ сервера
	headerContentTtype := r.Header.Get("Content-Type")
	// проверяем соответсвтвие типа содержимого запроса
	if headerContentTtype != "application/json" {
		msg = "Incorrect content type!"
		log.Print(msg)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(msg)
		return
	}
	// декодируем тело запроса
	var reqB AddMeetingRequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqB)
	if err != nil {
		msg = "Failed to decode request body: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m := database.Meeting{
		Name:                 reqB.Name,
		Description:          reqB.Description,
		ProfessorId:          reqB.ProfessorId,
		MeetingTime:          reqB.MeetingTime,
		StudentParticipantId: reqB.StudentId,
		IsOnline:             reqB.IsOnline,
	}
	h.app.Database.AddMeeting(m)
	w.WriteHeader(http.StatusOK)
}

// Добавляет пользователя в сегменты и удаляет пользователя из сегментов, указанных в теле запроса
func (h *Handler) GetMeeting(w http.ResponseWriter, r *http.Request) {
	var msg string      // ответ сервера
	vars := mux.Vars(r) //получаем переменные запроса

	meetId, err := strconv.Atoi(vars["meeting_id"])
	if err != nil {
		msg = "Unable to parse id: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	headerContentTtype := r.Header.Get("Content-Type")
	// проверяем соответсвтвие типа содержимого запроса
	if headerContentTtype != "application/json" {
		msg = "Incorrect content type!"
		log.Print(msg)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(msg)
		return
	}
	// декодируем тело запроса
	var reqB BaseRequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&reqB)
	if err != nil {
		msg = "Failed to decode request body: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := GetMeetingResultBody{
		Meeting: Meeting{},
	}
	mt, isFound := h.app.Database.GetMeeting(meetId)

	if !isFound {
		msg = "Failed to find meeting!"
		log.Print(msg)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	st, _ := h.app.Database.GetStudent(mt.StudentParticipantId)
	res.Meeting = Meeting{
		Id:          mt.Id,
		Name:        mt.Name,
		Description: mt.Description,
		MeetingTime: mt.MeetingTime,
		StudentParticipant: Student{
			Id:           st.Id,
			Name:         st.Name,
			Surname:      st.Surname,
			Middlename:   st.Middlename,
			Cource:       st.Cource,
			ProjectTheme: st.ProjectTheme,
		},
		IsOnline: mt.IsOnline}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
