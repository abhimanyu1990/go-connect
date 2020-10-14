package utility

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Referrer-Policy", "no-referrer")
	
	data["IsSuccess"] = true
	json.NewEncoder(w).Encode(data)
}
