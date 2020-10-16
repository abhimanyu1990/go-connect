package utility

import (
	"encoding/json"
	"net/http"
	"time"
	"go-connect/app/dto"
)

func RespondError(w http.ResponseWriter, message string, httpStatus int) {
	data := dto.ResponseError{time.Now().Unix(),false,message,httpStatus,time.Now()}
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}


func NotFoundError(w http.ResponseWriter, message string) {
	data := dto.ResponseError{time.Now().Unix(),false,message,404,time.Now()}
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(data)
}


func BadRequestError(w http.ResponseWriter, message string) {
	data := dto.ResponseError{time.Now().Unix(),false,message,400,time.Now()}
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(data)
}

func ForbiddenError(w http.ResponseWriter, message string) {
	data := dto.ResponseError{time.Now().Unix(),false,message,403,time.Now()}
	w.WriteHeader(403)
	json.NewEncoder(w).Encode(data)
}


