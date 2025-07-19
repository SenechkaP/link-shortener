package link

import (
	"advpractice/configs"
	"advpractice/pkg/event"
	"advpractice/pkg/middleware"
	"advpractice/pkg/req"
	"advpractice/pkg/res"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

const (
	ErrNonExistingLinkID   = "LINK WITH THIS ID IS NOT FOUND"
	ErrAlreadyExistingHash = "LINK WITH THIS HASH IS ALREADY EXISTING"
	ErrQueryParams         = "INVALID QUERY PARAMETERS"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
	EventBus       *event.EventBus
}

type LinkHandler struct {
	LinkRepository *LinkRepository
	EventBus       *event.EventBus
}

func NewLinkHandler(router *http.ServeMux, deps *LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
		EventBus:       deps.EventBus,
	}
	router.HandleFunc("GET /{hash}", handler.goTo())
	router.HandleFunc("POST /link", handler.createLink())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.updateLink(), deps.Config))
	router.HandleFunc("DELETE /link/{id}", handler.deleteLink())
	router.HandleFunc("GET /link", handler.getAllLinks())
}

func (handler *LinkHandler) goTo() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		hash := q.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusNotFound)
			return
		}
		go handler.EventBus.Publish(event.Event{
			Type: event.EventLinkVisited,
			Data: link.ID,
		})
		http.Redirect(w, q, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) createLink() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](q)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusBadRequest)
			return
		}
		link := NewLink(body.Url, handler.LinkRepository)
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.JsonDump(w, createdLink, http.StatusCreated)
	}
}

func (handler *LinkHandler) updateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		body, err := req.HandleBody[LinkUpdateRequest](q)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusBadRequest)
			return
		}

		userEmail := q.Context().Value(middleware.ContextEmailKey).(string)
		log.Println(userEmail)

		idString := q.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusBadRequest)
			return
		}

		if _, err := handler.LinkRepository.GetByID(uint(id)); err != nil {
			res.JsonDump(w, ErrNonExistingLinkID, http.StatusNotFound)
			return
		}
		if link, _ := handler.LinkRepository.GetByHash(body.Hash); link != nil {
			res.JsonDump(w, ErrAlreadyExistingHash, http.StatusBadRequest)
			return
		}

		updatedLink, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})

		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.JsonDump(w, updatedLink, http.StatusOK)
	}
}

func (handler *LinkHandler) deleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		idString := q.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := handler.LinkRepository.GetByID(uint(id)); err != nil {
			res.JsonDump(w, ErrNonExistingLinkID, http.StatusNotFound)
			return
		}
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (handler *LinkHandler) getAllLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		limit, errLimit := strconv.Atoi(q.URL.Query().Get("limit"))
		offset, errOffset := strconv.Atoi(q.URL.Query().Get("offset"))
		if errLimit != nil || errOffset != nil {
			res.JsonDump(w, ErrQueryParams, http.StatusBadRequest)
			return
		}
		count := handler.LinkRepository.Count()
		links := handler.LinkRepository.GetAll(limit, offset)
		res.JsonDump(w, AllLinksResponce{count, links}, http.StatusOK)
	}
}
