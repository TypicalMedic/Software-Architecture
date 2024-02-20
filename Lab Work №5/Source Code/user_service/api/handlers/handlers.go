package handlers

import (
	"PAPS-lab5/user_service/internal/app"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	req, err := http.NewRequest("GET", "http://meeting-service-container:8081/meetings/filter", nil)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	q := req.URL.Query()
	q.Add("student", vars["student"])
	q.Add("prof", fmt.Sprint(reqB.ProfessorId))
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		msg = "Failed call to meetings service: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var meetResp GetMeetingsMeetingServiceResultBody
	decoder = json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&meetResp)
	if err != nil {
		msg = "Failed to decode meeting responce body: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := GetMeetingsResultBody{
		Meetings: []Meeting{},
	}
	for _, mt := range meetResp.Meetings {
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
	req, err := http.NewRequest("GET", "http://meeting-service-container:8081/meetings", nil)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	q := req.URL.Query()
	q.Add("prof", fmt.Sprint(reqB.ProfessorId))
	log.Print(reqB.ProfessorId)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		msg = "Failed call to meetings service: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var meetResp GetMeetingsMeetingServiceResultBody
	decoder = json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&meetResp)
	if err != nil {
		msg = "Failed to decode meeting responce body: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := GetMeetingsResultBody{
		Meetings: []Meeting{},
	}
	for _, mt := range meetResp.Meetings {
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

	resp, err := http.Post("http://meeting-service-container:8081/meeting/add", "application/json", r.Body)
	if err != nil {
		msg = "Failed call to meetings service: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(resp.StatusCode)
}

// Добавляет пользователя в сегменты и удаляет пользователя из сегментов, указанных в теле запроса
func (h *Handler) GetMeeting(w http.ResponseWriter, r *http.Request) {
	var msg string      // ответ сервера
	vars := mux.Vars(r) //получаем переменные запроса

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

	req, err := http.NewRequest("GET", "http://meeting-service-container:8081/meeting/"+vars["meeting_id"], nil)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	q := req.URL.Query()
	q.Add("prof", fmt.Sprint(reqB.ProfessorId))
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		msg = "Failed call to meetings service: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if resp.StatusCode != 200 {
		log.Print("meetings server returned bad status: " + resp.Status)
		w.WriteHeader(resp.StatusCode)
		return
	}

	var meetResp GetMeetingMeetingServiceResultBody
	decoder = json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&meetResp)
	if err != nil {
		msg = "Failed to decode meeting responce body: " + err.Error()
		log.Print(msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := GetMeetingResultBody{
		Meeting: Meeting{},
	}

	st, _ := h.app.Database.GetStudent(meetResp.Meeting.StudentParticipantId)
	res.Meeting = Meeting{
		Id:          meetResp.Meeting.Id,
		Name:        meetResp.Meeting.Name,
		Description: meetResp.Meeting.Description,
		MeetingTime: meetResp.Meeting.MeetingTime,
		StudentParticipant: Student{
			Id:           st.Id,
			Name:         st.Name,
			Surname:      st.Surname,
			Middlename:   st.Middlename,
			Cource:       st.Cource,
			ProjectTheme: st.ProjectTheme,
		},
		IsOnline: meetResp.Meeting.IsOnline}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
