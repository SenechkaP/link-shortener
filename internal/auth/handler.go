package auth

import (
	"advpractice/configs"
	"advpractice/pkg/req"
	"advpractice/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{deps.Config}
	router.HandleFunc("POST /auth/login", handler.login())
	router.HandleFunc("POST /auth/register", handler.register())
}

func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {

		body, err := req.HandleBody[LoginRequest](q)
		if err != nil {
			res.JsonDump(w, err.Error(), 402)
			return
		}
		res.JsonDump(w, body, 201)
	}
}

func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		body, err := req.HandleBody[RegistrateRequest](q)
		if err != nil {
			res.JsonDump(w, err.Error(), 402)
			return
		}
		res.JsonDump(w, body, 201)
	}
}
