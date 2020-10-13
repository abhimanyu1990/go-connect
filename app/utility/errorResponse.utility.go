package utility

import (
	"encoding/json"
	"net/http"
	"time"
)

func RespondError(w http.ResponseWriter, message string, httpStatus int) {
	data := map[string]interface{}{
									 "isSuccess": false,
									 "message": message,
									 "time":time.Now(),
									 "errorId":time.Now().Unix(),
									 "responseCode":httpStatus,
								}
	w.Header().Add("Content-Type","application/json")	
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}

