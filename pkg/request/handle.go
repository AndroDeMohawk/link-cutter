package request

import (
	"net/http"

	"github.com/AndroDeMohawk/link-cutter/pkg/response"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	var payload T
	body, err := Decode[T](r.Body)
	if err != nil {
		response.Send_json(*w, 402, payload)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		response.Send_json(*w, 402, payload)
		return nil, err
	}
	return &body, err
}
