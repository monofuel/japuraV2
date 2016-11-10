package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//AppHandler is a wrapper around api requests that returns an AppError
type AppHandler func(http.ResponseWriter, *http.Request) *AppError

//AppError is a wrapper around useful information to return in an error
type AppError struct {
	Error   error
	Message string
	Code    int
}

//NewAppError returns a shiny new AppError
func NewAppError(err error, message string, code int) *AppError {
	return &AppError{
		Error:   err,
		Message: message,
		Code:    code,
	}
}

//ServeHTTP implements the http handler for AppHandlers
func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var errorResponse struct {
		Err     error  `json:"err"`
		Message string `json:"message"`
		Code    int    `json:"status"`
	}
	if e := fn(w, r); e != nil {
		fmt.Printf("error: %v, %v", e.Message, e.Error)
		//http.Error(w, e.Message, e.Code)
		errorResponse.Err = e.Error
		errorResponse.Message = e.Message
		errorResponse.Code = e.Code
		w.WriteHeader(e.Code)
		json.NewEncoder(w).Encode(errorResponse)
	}
}

//GetPageLimitParameters is a handy function to get the page and limit parameters
func GetPageLimitParameters(r *http.Request) (int, int, error) {
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return 0, 0, err
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return 0, 0, err
	}
	return page, limit, nil
}

func UnixTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
