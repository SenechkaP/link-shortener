package auth

import (
	"advpractice/configs"
	"advpractice/pkg/jwt"
	"advpractice/pkg/req"
	"advpractice/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{deps.Config, deps.AuthService}
	router.HandleFunc("POST /auth/login", handler.login())
	router.HandleFunc("POST /auth/register", handler.register())
}

func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		body, err := req.HandleBody[LoginRequest](q)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.AuthService.login(body)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := handler.AuthService.getJWT(jwt.JWTData{Email: body.Email}, handler.Config.Auth.Secret)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.JsonDump(w, TokenResponse{token}, http.StatusOK)
	}
}

func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		body, err := req.HandleBody[RegistrateRequest](q)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.AuthService.register(body)
		if err != nil {
			if err.Error() == ErrUserExists {
				res.JsonDump(w, err.Error(), http.StatusConflict)
				return
			}
			if err.Error() == ErrWrongPassword {
				res.JsonDump(w, err.Error(), http.StatusBadRequest)
				return
			}
			res.JsonDump(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := handler.AuthService.getJWT(jwt.JWTData{Email: body.Email}, handler.Config.Auth.Secret)
		if err != nil {
			res.JsonDump(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.JsonDump(w, TokenResponse{token}, http.StatusOK)
	}
}
