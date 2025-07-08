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
	router.HandleFunc("GET /{hash}", handler.getLink())
	router.HandleFunc("POST /link", handler.createLink())
	router.HandleFunc("PATCH /link/{id}", handler.patchLink())
	router.HandleFunc("DELETE /link/{id}", handler.deleteLink())

}

func (handler *LinkHandler) getLink() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
	}
}

func (handler *LinkHandler) createLink() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](q)
		if err != nil {
			res.JsonDump(w, err.Error(), 402)
			return
		}
		link := NewLink(body.Url)
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			res.JsonDump(w, err.Error(), 500)
		}
		res.JsonDump(w, createdLink, 201)
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
