package database

// интерфейс базы данных, где прописаны методы взаимодействия с ней
type IDatabase interface {
	// проверка работы бд
	Ping() error
	GetProf(profId int) (Professor, bool)
	GetStudent(studId int) (Student, bool)
}
