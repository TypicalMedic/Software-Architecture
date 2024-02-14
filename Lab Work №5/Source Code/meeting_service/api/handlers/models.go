package handlers

import "time"

// структуры для json тел запросов и ответов

type BaseRequestBody struct {
	ProfessorId int `json:"professor_id"`
}

type AddMeetingRequestBody struct {
	BaseRequestBody
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MeetingTime time.Time `json:"meeting_time"`
	StudentId   int       `json:"student_participant_id"`
	IsOnline    bool      `json:"is_online"`
}

type GetMeetingsResultBody struct {
	Meetings []Meeting `json:"meetings"`
}

type GetMeetingResultBody struct {
	Meeting Meeting `json:"meeting"`
}
type Student struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Middlename   string `json:"middlename"`
	Cource       int    `json:"cource"`
	ProjectTheme string `json:"project_theme"`
}

type Meeting struct {
	Id                   int       `json:"id"`
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	MeetingTime          time.Time `json:"meeting_time"`
	StudentParticipantId int       `json:"student_participant_id"`
	IsOnline             bool      `json:"is_online"`
}
