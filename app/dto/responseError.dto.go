package dto

import(
	"time"
)

type ResponseError struct {
		ErrorId int64
		IsSuccess bool
		Message string
		ResponseCode int
		TimeOfDay time.Time
}