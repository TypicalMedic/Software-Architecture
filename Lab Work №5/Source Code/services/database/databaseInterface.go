package database

// интерфейс базы данных, где прописаны методы взаимодействия с ней
type IDatabase interface {
	// проверка работы бд
	Ping() error

	GetProfMeetings(profId int) []Meeting
	GetMeeting(meetId int) (Meeting, bool)
	GetProfsStudentMeetings(profId, studId int) []Meeting
	GetProf(profId int) (Professor, bool)
	GetStudent(studId int) (Student, bool)
	AddMeeting(m Meeting)
}
