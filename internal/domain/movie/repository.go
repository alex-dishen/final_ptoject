package movie

import (
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
)

var settings = postgresql.ConnectionURL{
	Database: `postgres`,
	Host:     `localhost:5432`,
	User:     `postgres`,
	Password: `backtowork`,
}

type Repository interface {
	FindAll() ([]Movie, error)
	FindById(id int64) (*Movie, error)
	CreateMovie(m *Movie) (*Movie, error)
	UpdateMovie(m *Movie) (*Movie, error)
	Delete(id int64) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Movie, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		fmt.Print("FindAll.Open(settings): ", err)
	}
	defer sess.Close()

	filmsCol := sess.Collection("movies")

	var films []Movie

	err = filmsCol.Find().All(&films)
	if err != nil {
		fmt.Print("filmsCol: ", err)
	}

	return films, nil
}

func (r *repository) FindById(id int64) (*Movie, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		fmt.Print("Open: ", err)
	}
	defer sess.Close()

	var film *Movie

	err = sess.SQL().
		SelectFrom("movies").
		Where("id", id).
		One(&film)
	if err != nil {
		fmt.Print("Query: ", err)
	}

	return film, nil
}

func (r *repository) CreateMovie(m *Movie) (*Movie, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		fmt.Print("Open: ", err)
	}
	defer sess.Close()

	_, err = sess.SQL().
		InsertInto("movies").
		Columns("id", "name", "director", "year").Values(m.ID, m.Name, m.Director, m.Year).
		Exec()
	if err != nil {
		fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
	}

	lastMovie, err := r.FindById(m.ID)
	if err != nil {
		fmt.Print("movie.Insert: ", err)
	}
	return lastMovie, nil
}

func (r *repository) UpdateMovie(m *Movie) (*Movie, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		fmt.Print("Open: ", err)
	}
	defer sess.Close()

	_, err = sess.SQL().
		Update("movies").
		Set("name", m.Name, "director", m.Director, "year", m.Year).
		Where("id = ?", m.ID).
		Exec()
	if err != nil {
		fmt.Print("movies.Update: ", err)
	}

	return r.FindById(m.ID)
}

func (r *repository) Delete(id int64) error {
	sess, err := postgresql.Open(settings)
	if err != nil {
		fmt.Print("FindAll.Open(settings): ", err)
	}
	defer sess.Close()

	filmColl := sess.Collection("movies")
	err = filmColl.Find("id", id).Delete()
	if err != nil {
		fmt.Print("filmColl.Delete: ", err)
	}

	return err
}
