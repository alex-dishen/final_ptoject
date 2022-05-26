package controllers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/event"
	"net/http"
	"strconv"
)

type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		event, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, event)
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

		event, err := (*c.service).CreateMovie(name, director, year)
		if err != nil {
			fmt.Printf("EventController.CreateMovie(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.CreateMovie(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.CreateMovie(): %s", err)
		}
	}
}

func (c *EventController) UpdateName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.CreateMovie(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.CreateMovie(): %s", err)
			}
			return
		}
		name := chi.URLParam(r, "name")
		director := chi.URLParam(r, "director")
		year, err := strconv.ParseInt(chi.URLParam(r, "year"), 10, 64)

		event, err := (*c.service).UpdateName(id, name, director, year)
		if err != nil {
			fmt.Printf("EventController.CreateMovie(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.CreateMovie(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.CreateMovie(): %s", err)
		}
	}
}
