package controllers

import (
	"encoding/json"
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

		var film *movie.Movie

		err := json.NewDecoder(r.Body).Decode(&film)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Create(): %s", err)
			}
			return
		}

		movies, err := (*c.service).CreateMovie(film)
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

		var film *movie.Movie

		err := json.NewDecoder(r.Body).Decode(&film)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
			}
			return
		}
		movies, err := (*c.service).UpdateMovie(film)
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

func (c *EventController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("MovieController.Delete(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("MovieController.Delete(): %s", err)
			}
			return
		}
		res := (*c.service).Delete(id)

		if err != nil {
			fmt.Printf("MovieController.Delete(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("MovieController.Delete(): %s", err)
			}
			return
		}

		err = success(w, res)
		if err != nil {
			fmt.Printf("MovieController.Delete(): %s", err)
		}
	}
}
