package controllers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/movie"
	"net/http"
	"strconv"
)

type EventController struct {
	service *movie.Service
}

func NewEventController(s *movie.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, movies)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindById(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindById(): %s", err)
			}
			return
		}
		movies, err := (*c.service).FindById(id)
		if err != nil {
			fmt.Printf("EventController.FindById(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindById(): %s", err)
			}
			return
		}

		err = success(w, movies)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

func (c *EventController) CreateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		director := chi.URLParam(r, "director")
		year, err := strconv.ParseInt(chi.URLParam(r, "year"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.CreateMovie(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.CreateMovie(): %s", err)
			}
			return
		}

		movies, err := (*c.service).CreateMovie(name, director, year)
		if err != nil {
			fmt.Printf("EventController.CreateMovie(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.CreateMovie(): %s", err)
			}
			return
		}

		err = success(w, movies)
		if err != nil {
			fmt.Printf("EventController.CreateMovie(): %s", err)
		}
	}
}

func (c *EventController) UpdateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.UpdateMovie(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.UpdateMovie(): %s", err)
			}
			return
		}
		name := chi.URLParam(r, "name")
		director := chi.URLParam(r, "director")
		year, err := strconv.ParseInt(chi.URLParam(r, "year"), 10, 64)

		movies, err := (*c.service).UpdateMovie(id, name, director, year)
		if err != nil {
			fmt.Printf("EventController.UpdateMovie(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.UpdateMovie(): %s", err)
			}
			return
		}

		err = success(w, movies)
		if err != nil {
			fmt.Printf("EventController.UpdateMovie(): %s", err)
		}
	}
}
