package database

// интерфейс базы данных, где прописаны методы взаимодействия с ней
type IDatabase interface {
	// проверка работы бд
	Ping() error

	GetProfMeetings(profId int) []Meeting
	GetMeeting(meetId int) (Meeting, bool)
	GetProfsStudentMeetings(profId, studId int) []Meeting
	AddMeeting(m Meeting)
}
