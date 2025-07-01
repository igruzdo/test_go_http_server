package auth

import (
	"fmt"
	"http_server/configs"
	"http_server/pakages/response"
	"net/http"
)

type AuthHandlerDeps struct{
	*configs.Config
}
 
type AuthHandler struct{
	*configs.Config
}


func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps)  { 
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func (writer http.ResponseWriter, resp *http.Request)  {
		fmt.Println(handler.Config.Auth.Secret)
		res := LoginResponse{
			Token: "aaaa",
		}

		response.Json(writer, res, 201 )
	}
} 
 
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func (writer http.ResponseWriter, resp *http.Request)  {
		fmt.Println("REGISTER")
	} 
}