package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// реализация интерфейса баз данных, содержит ссылку на подключенную бд и имеет методы реализуемого интерфеса
type Database struct {
	db *sql.DB
}

// конструктор новой бд
func NewDatabase(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

// проверка работы бд
func (db *Database) Ping() error {
	return db.db.Ping()
}

func (db *Database) GetProfMeetings(profId int) []Meeting {
	meeetings := []Meeting{}
	var meeting Meeting
	sql := fmt.Sprint("SELECT * FROM meeting WHERE professor_id = ", profId, ";")
	res, err := db.db.Query(sql)
	if err != nil {
		log.Fatal(err)
		return []Meeting{}
	}

	for res.Next() {
		if err = res.Scan(&meeting.Id, &meeting.Name, &meeting.Description, &meeting.MeetingTime, &meeting.StudentParticipantId, &meeting.IsOnline, &meeting.ProfessorId); err != nil {
			log.Fatal(err)
			return []Meeting{}
		}
		// добавляем в массив
		meeetings = append(meeetings, meeting)
	}
	return meeetings
}
func (db *Database) GetMeeting(meetId int) (Meeting, bool) {
	var meeting Meeting
	sql := fmt.Sprint("SELECT * FROM meeting WHERE id = ", meetId, ";")
	res := db.db.QueryRow(sql)

	err := res.Scan(&meeting.Id, &meeting.Name, &meeting.Description, &meeting.MeetingTime, &meeting.StudentParticipantId, &meeting.IsOnline, &meeting.ProfessorId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return Meeting{}, false
		}
		log.Fatal(err)
		return Meeting{}, false
	}
	return meeting, true

}
func (db *Database) GetProfsStudentMeetings(profId, studId int) []Meeting {
	meeetings := []Meeting{}
	var meeting Meeting
	sql := fmt.Sprint("SELECT * FROM meeting WHERE professor_id = ", profId, " and student_participant_id = ", studId, ";")
	res, err := db.db.Query(sql)
	if err != nil {
		log.Fatal(err)
		return []Meeting{}
	}

	for res.Next() {
		if err = res.Scan(&meeting.Id, &meeting.Name, &meeting.Description, &meeting.MeetingTime, &meeting.StudentParticipantId, &meeting.IsOnline, &meeting.ProfessorId); err != nil {
			log.Fatal(err)
			return []Meeting{}
		}
		// добавляем в массив
		meeetings = append(meeetings, meeting)
	}
	return meeetings
}

func (db *Database) GetProf(profId int) (Professor, bool) {
	var professor Professor
	sql := fmt.Sprint("SELECT * FROM professor WHERE id = ", profId, ";")
	res := db.db.QueryRow(sql)

	err := res.Scan(&professor.Id, &professor.Name, &professor.Surname, &professor.Middlename, &professor.CalendarEmail)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return Professor{}, false
		}
		log.Fatal(err)
		return Professor{}, false
	}
	return professor, true
}

func (db *Database) GetStudent(studId int) (Student, bool) {
	var student Student
	sql := fmt.Sprint("SELECT * FROM student WHERE id = ", studId, ";")
	res := db.db.QueryRow(sql)

	err := res.Scan(&student.Id, &student.Name, &student.Surname, &student.Middlename, &student.Cource, &student.ProjectTheme)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return Student{}, false
		}
		log.Fatal(err)
		return Student{}, false
	}
	return student, true
}

func (db *Database) AddMeeting(m Meeting) {
	sql := "INSERT INTO meeting (name, description, meeting_time, student_participant_id, is_online, professor_id) VALUES (?, ?, ?, ?, ?, ?);"
	_, err := db.db.Exec(sql, m.Name, m.Description, m.MeetingTime, m.StudentParticipantId, m.IsOnline, m.ProfessorId)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// Создает соединение с БД по строке подключения. Проверяет успешность и выводит ошибки, если такие есть.
//
// Возвращает подключенную БД при успешном подкючении, иначе nil
func CreateDatabaseConnectionMySQL(DatabaseConnectionString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", DatabaseConnectionString)
	if err == nil {
		// проверяем, подключились ли к бд
		err := db.Ping()
		if err != nil {
			return nil, err
		}
		return db, nil
	} else {
		return nil, err
	}
}
