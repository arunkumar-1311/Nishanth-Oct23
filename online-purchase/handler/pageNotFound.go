package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handlers) PageNotFound(w http.ResponseWriter, r *http.Request) {

	message := map[string]interface{}{
		"error":   fmt.Sprintf("%v page not found", r.URL.Path),
		"message": "Enter a valid path ,For reference user ReadMe.md",
	}
	json.NewEncoder(w).Encode(message)
}
