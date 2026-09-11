package response

import (
	"encoding/json"
	"net/http"
)

func Send_json(w http.ResponseWriter, Status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(Status)
	json.NewEncoder(w).Encode(body)
}
