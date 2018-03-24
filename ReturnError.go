package restful

import (
	"encoding/json"
	"net/http"
)

func ReturnError(w http.ResponseWriter, status int, error HttpError) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}
