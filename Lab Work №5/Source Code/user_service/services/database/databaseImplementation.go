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
