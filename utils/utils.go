package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON mengirim response dalam format JSON
func RespondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}