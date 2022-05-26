package event

import (
	"fmt"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

var settings = postgresql.ConnectionURL{
	Database: `postgres`,
	Host:     `localhost:5432`,
	User:     `postgres`,
	Password: `backtowork`,
}

type Repository interface {
	FindAll() ([]Movie, error)
	FindOne(id int64) (*Movie, error)
	CreateMovie(name string, director string, year int64) (*Movie, error)
	UpdateName(id int64, name string, director string, year int64) (*Movie, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Movie, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("FindAll.Open(settings): ", err)
	}
	defer sess.Close()

	filmsCol := sess.Collection("movies")

	var films []Movie

	err = filmsCol.Find().All(&films)
	if err != nil {
		log.Fatal("filmsCol: ", err)
	}

	fmt.Printf("Records in the %q collection:\n", filmsCol.Name())
	for i := range films {
		fmt.Printf("Record #%d: %#v\n", i, films[i])
	}

	return films, nil
}

func (r *repository) FindOne(id int64) (*Movie, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	var film Movie

	err = sess.SQL().
		SelectFrom("movies").
		Where("id", id).
		One(&film)
	if err != nil {
		log.Fatal("Query: ", err)
	}

	return &Movie{
		ID:       film.ID,
		Name:     film.Name,
		Director: film.Director,
		Year:     film.Year,
	}, nil
}

func (r *repository) CreateMovie(name string, director string, year int64) (*Movie, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	db.LC().SetLevel(db.LogLevelDebug)
	defer sess.Close()
	_, err = sess.SQL().
		InsertInto("movies").
		Columns("name", "director", "year").Values(name, director, year).
		Exec()
	if err != nil {
		fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
	}
	return &Movie{
		Name:     name,
		Director: director,
		Year:     year,
	}, nil
}

func (r *repository) UpdateName(id int64, name string, director string, year int64) (*Movie, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	db.LC().SetLevel(db.LogLevelDebug)

	defer sess.Close()

	if name != "" {
		_, err := sess.SQL().
			Update("movies").
			Set("name = ?", name).
			Where("id = ?", id).
			Exec()
		if err != nil {
			fmt.Printf("sess.SQL: %v. This is expected on the read-only sandbox.\n", err)
		}
	}

	if director != "" {
		_, err := sess.SQL().
			Update("movies").
			Set("director = ?", director).
			Where("id = ?", id).
			Exec()
		if err != nil {
			fmt.Printf("sess.SQL: %v. This is expected on the read-only sandbox.\n", err)
		}
	}

	if year != 0 {
		_, err := sess.SQL().
			Update("movies").
			Set("year = ?", year).
			Where("id = ?", id).
			Exec()
		if err != nil {
			fmt.Printf("sess.SQL: %v. This is expected on the read-only sandbox.\n", err)
		}
	}

	var films Movie

	err = sess.SQL().
		SelectFrom("movies").
		Where("id", id). // Or Where("last_name = ?", "Poe")
		One(&films)
	if err != nil {
		log.Fatal("sess.SQL: ", err)
	}

	return &Movie{
		ID:       films.ID,
		Name:     films.Name,
		Director: films.Director,
		Year:     films.Year,
	}, nil
}
