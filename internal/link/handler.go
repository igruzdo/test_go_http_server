package link

import (
	"http_server/pakages/request"
	"http_server/pakages/response"
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
		body, err := request.HandleBody[LinkCreateRequest](&writer, req)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(writer, createdLink, 201)
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
		hash := req.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
		}

		http.Redirect(writer, req, link.Url, http.StatusTemporaryRedirect)
	}
}
