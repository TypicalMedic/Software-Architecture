package handlers

import (
	"PAPS-lab5/meeting_service/internal/app"
	"PAPS-lab5/meeting_service/services/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	var msg string // ответ сервера
	msg = fmt.Sprint(h.app.Configuration.Server.Host, ": GET /meetings/filter")
	log.Print(msg)
	studId, err := strconv.Atoi(r.URL.Query().Get("student"))
	if err != nil {
		msg = "Unable to parse id: " + err.Error()
		log.Print(msg)
		msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusBadRequest)
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	prof, err := strconv.Atoi(r.URL.Query().Get("prof"))
	if err != nil {
		msg = "Unable to parse id: " + err.Error()
		log.Print(msg)
		msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusBadRequest)
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := GetMeetingsResultBody{
		Meetings: []Meeting{},
	}
	m := h.app.Database.GetProfsStudentMeetings(prof, studId)
	for _, mt := range m {
		res.Meetings = append(res.Meetings,
			Meeting{
				Id:                   mt.Id,
				Name:                 mt.Name,
				Description:          mt.Description,
				MeetingTime:          mt.MeetingTime,
				StudentParticipantId: mt.StudentParticipantId,
				IsOnline:             mt.IsOnline})
	}
	msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusOK)
	log.Print(msg)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Получает список сегментов, в которых находится пользователь по его id
func (h *Handler) GetProfMeetings(w http.ResponseWriter, r *http.Request) {
	var msg string // ответ сервера
	msg = fmt.Sprint(h.app.Configuration.Server.Host, ": GET /meetings")
	log.Print(msg)
	prof, err := strconv.Atoi(r.URL.Query().Get("prof"))
	if err != nil {
		msg = "Unable to parse id: " + err.Error()
		log.Print(msg)
		msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusBadRequest)
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := GetMeetingsResultBody{
		Meetings: []Meeting{},
	}
	m := h.app.Database.GetProfMeetings(prof)
	for _, mt := range m {
		res.Meetings = append(res.Meetings,
			Meeting{
				Id:                   mt.Id,
				Name:                 mt.Name,
				Description:          mt.Description,
				MeetingTime:          mt.MeetingTime,
				StudentParticipantId: mt.StudentParticipantId,
				IsOnline:             mt.IsOnline})
	}
	msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusOK)
	log.Print(msg)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Добавляет пользователя в бд
func (h *Handler) AddMeeting(w http.ResponseWriter, r *http.Request) {

	var msg string // ответ сервера
	msg = fmt.Sprint(h.app.Configuration.Server.Host, ": POST /meeting/add")
	log.Print(msg)
	headerContentTtype := r.Header.Get("Content-Type")
	// проверяем соответсвтвие типа содержимого запроса
	if headerContentTtype != "application/json" {
		msg = "Incorrect content type!"
		log.Print(msg)
		msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusUnsupportedMediaType)
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
		msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusBadRequest)
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
	msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusOK)
	log.Print(msg)
	w.WriteHeader(http.StatusOK)
}

// Добавляет пользователя в сегменты и удаляет пользователя из сегментов, указанных в теле запроса
func (h *Handler) GetMeeting(w http.ResponseWriter, r *http.Request) {
	var msg string // ответ сервера
	msg = fmt.Sprint(h.app.Configuration.Server.Host, ": GET /meeting/{id}")
	log.Print(msg)

	meetId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/meeting/"))
	if err != nil {
		msg = "Unable to parse id: " + err.Error()
		log.Print(msg)
		msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusBadRequest)
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	prof, err := strconv.Atoi(r.URL.Query().Get("prof"))
	if err != nil {
		msg = "Unable to parse id: " + err.Error()
		log.Print(msg)
		msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusBadRequest)
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
		msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusNotFound)
		log.Print(msg)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if mt.ProfessorId != prof {
		msg = "Requested meeting is not current professor's!"
		log.Print(msg)
		msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusForbidden)
		log.Print(msg)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	res.Meeting = Meeting{
		Id:                   mt.Id,
		Name:                 mt.Name,
		Description:          mt.Description,
		MeetingTime:          mt.MeetingTime,
		StudentParticipantId: mt.StudentParticipantId,
		IsOnline:             mt.IsOnline}

	msg = fmt.Sprint(h.app.Configuration.Server.Host, ": return ", http.StatusOK)
	log.Print(msg)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
