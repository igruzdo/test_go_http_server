package link

import (
	"fmt"
	"http_server/configs"
	"http_server/pakages/event"
	"http_server/pakages/middleware"
	"http_server/pakages/request"
	"http_server/pakages/response"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
	EventBus       *event.EventBus
}

type LinkHandler struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
	EventBus       *event.EventBus
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {

	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuth(handler.Update(), deps.Config))
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
		for {

			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}

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

		email, ok := req.Context().Value(middleware.ContextEmailKey).(string)

		if ok {
			fmt.Println(email)
		}

		body, err := request.HandleBody[LinkUpdateRequest](&writer, req)
		if err != nil {
			return
		}

		idString := req.PathValue("id")
		id, err := strconv.ParseInt(idString, 10, 32)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		response.Json(writer, link, http.StatusOK)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		idString := req.PathValue("id")
		id, err := strconv.ParseInt(idString, 10, 32)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = handler.LinkRepository.GetById(uint(id))

		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		}

		err = handler.LinkRepository.Delete(uint(id))

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Json(writer, nil, 200)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusNotFound)
		}
		//handler.StatRepository.AddClick(link.ID)
		handler.EventBus.Publish(event.Event{
			Type: event.LinkVisitedEvent,
			Data: link.ID,
		})
		http.Redirect(writer, req, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) GetAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		limit, err := strconv.Atoi(req.URL.Query().Get("limit"))

		if err != nil {
			http.Error(writer, "invalid limit", http.StatusBadRequest)
			return
		}

		offset, err := strconv.Atoi(req.URL.Query().Get("offset"))

		if err != nil {
			http.Error(writer, "invalid offset", http.StatusBadRequest)
			return
		}

		links := handler.LinkRepository.GetLinks(uint(limit), uint(offset))

		response.Json(writer, links, http.StatusOK)
	}
}
