package repositories

import (
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Profile struct {
	Login      string
	Password   string
	Name       string
	Surname    string
	Age        int32
	Gender     bool
	HasAnyPets bool
	Employed   bool
}

func OpenPostgreSQLRepository() (Repository, error) {
	db, err := sql.Open("postgres", "su:su@/profiles")
	if err != nil {
		log.Println(err)
	}
	return &ProfileRepository{db: db}, nil
}

type Repository interface {
	InsertIntoDB(user Profile) error
	CloseDB() error
}
type ProfileRepository struct {
	db *sql.DB
}

func (r ProfileRepository) InsertIntoDB(user Profile) error {
	_, err := r.db.Query("INSERT INTO users VALUES (?, ?, ?, ?, ?, ?, ?, ?)", user.Login, user.Name, user.Surname,
		user.Password, user.Gender, user.HasAnyPets, user.Employed, user.Age)
	if err != nil {
		log.Errorf("There are some problems during insertion in table")
		log.Println(err)
	}
	return err
}
func (r ProfileRepository) CloseDB() error {
	r.db.Close()
	log.Info("Database is closed")
	return nil
}
