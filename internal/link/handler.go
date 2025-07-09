package link

import (
	"advpractice/pkg/req"
	"advpractice/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps *LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("GET /{hash}", handler.goTo())
	router.HandleFunc("POST /link", handler.createLink())
	router.HandleFunc("PATCH /link/{id}", handler.updateLink())
	router.HandleFunc("DELETE /link/{id}", handler.deleteLink())

}

func (handler *LinkHandler) goTo() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		hash := q.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusNotFound)
			return
		}
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
		idString := q.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := handler.LinkRepository.GetByID(uint(id)); err != nil {
			res.JsonDump(w, "Link with this ID is not found", http.StatusNotFound)
			return
		}
		if link, _ := handler.LinkRepository.GetByHash(body.Hash); link != nil {
			res.JsonDump(w, "Link with this hash is already existing", http.StatusBadRequest)
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
			res.JsonDump(w, "Link with this ID is not found", http.StatusNotFound)
			return
		}
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
