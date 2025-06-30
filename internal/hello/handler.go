package hello

import (
	"fmt"
	"net/http"
)

type HalloHandler struct{}

func NewHalloHandler(router *http.ServeMux )  { 
	handler := &HalloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}

func (handler *HalloHandler) Hello() http.HandlerFunc {
	return func (writer http.ResponseWriter, resp *http.Request)  {
		fmt.Println("YES")
	}
} 