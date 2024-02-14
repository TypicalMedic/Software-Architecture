package database

import "time"

type Professor struct {
	Id            int
	Name          string
	Surname       string
	Middlename    string
	CalendarEmail string
}

type Student struct {
	Id           int
	Name         string
	Surname      string
	Middlename   string
	Cource       int
	ProjectTheme string
}

type Meeting struct {
	Id                   int
	Name                 string
	Description          string
	MeetingTime          time.Time
	StudentParticipantId int
	IsOnline             bool
	ProfessorId          int
}
