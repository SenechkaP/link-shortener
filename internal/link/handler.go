package link

import "net/http"

type LinkHandler struct{}

func NewLinkHabdler(router *http.ServeMux) {
	handler := &LinkHandler{}
	router.HandleFunc("GET /{alias}", handler.getLink())
	router.HandleFunc("POST /link", handler.postLink())
	router.HandleFunc("PATCH /link/{id}", handler.patchLink())
	router.HandleFunc("DELETE /link/{id}", handler.deleteLink())

}

func (handler *LinkHandler) getLink() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {

	}
}

func (handler *LinkHandler) postLink() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {

	}
}

func (handler *LinkHandler) patchLink() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {

	}
}

func (handler *LinkHandler) deleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {

	}
}
