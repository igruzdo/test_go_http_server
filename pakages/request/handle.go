package request

import (
	"http_server/pakages/response"
	"net/http"
)

func HandleBody[T any](writer *http.ResponseWriter, request *http.Request) (*T, error) {

	body, err := Decode[T](request.Body)

	if err != nil {
		response.Json(*writer, "json не декодирован", 402)
		return nil, err
	}

	err = IsValid[T](body)

	if err != nil {
		response.Json(*writer, "не соответствие схеме", 402)
		return nil, err
	}
	response.Json(*writer, body, 201)
	return &body, nil
}
