package service_errors

import "net/http"

type ServiceError struct {
	Code    int
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}

var (
	EmailExist = &ServiceError{http.StatusBadRequest, "email already registered"}
)
