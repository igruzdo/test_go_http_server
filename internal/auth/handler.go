package auth

import (
	"http_server/configs"
	"http_server/pakages/jwt"
	"http_server/pakages/request"
	"http_server/pakages/response"
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

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {

	return func(writer http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LoginRequest](&writer, req)

		userEmail, err := handler.AuthService.Login(body.Email, body.Password)

		if err != nil {
			response.Json(writer, err.Error(), http.StatusUnauthorized)
			return
		}

		jwtData := &jwt.JWTData{
			Email: userEmail,
		}

		token, err := jwt.NewJWT(handler.Auth.Secret).Create(jwtData)

		if err != nil {
			response.Json(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		res := LoginResponse{
			Token: token,
		}

		response.Json(writer, res, 201)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&writer, req)
		if err != nil {
			response.Json(writer, err.Error(), 201)
		}

		userEmail, err := handler.AuthService.Register(body.Email, body.Password, body.Name)

		if err != nil {
			response.Json(writer, err.Error(), http.StatusUnauthorized)
			return
		}

		jwtData := &jwt.JWTData{
			Email: userEmail,
		}

		token, err := jwt.NewJWT(handler.Auth.Secret).Create(jwtData)

		if err != nil {
			response.Json(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		res := LoginResponse{
			Token: token,
		}

		response.Json(writer, res, 201)
	}
}
