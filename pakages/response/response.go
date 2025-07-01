package response

import (
	"encoding/json"
	"net/http"
)

 func Json(writer http.ResponseWriter, res any, statusCode int) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(statusCode)
		json.NewEncoder(writer).Encode(res)
 }