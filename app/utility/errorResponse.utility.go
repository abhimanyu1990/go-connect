package utility

import (
	"encoding/json"
	"net/http"
	"time"
	"go-connect/app/dto"
)

func RespondError(w http.ResponseWriter, message string, httpStatus int) {
	data := dto.ResponseError{time.Now().Unix(),false,message,httpStatus,time.Now()}
	w.Header().Set("Content-Type", "application/json")	
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}

