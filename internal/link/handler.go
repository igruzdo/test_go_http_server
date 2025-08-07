package link

import (
	"net/http"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {

	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{alias}", handler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {

	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {

	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {

	}
}
