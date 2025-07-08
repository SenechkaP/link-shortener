package link

import (
	"advpractice/pkg/req"
	"advpractice/pkg/res"
	"fmt"
	"net/http"
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
	router.HandleFunc("PATCH /link/{id}", handler.patchLink())
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

func (handler *LinkHandler) patchLink() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		id := q.PathValue("id")
		fmt.Println(id)
	}
}

func (handler *LinkHandler) deleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {

	}
}
