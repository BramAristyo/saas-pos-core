package service_errors

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
)

type ServiceError struct {
	Code    int
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func IsUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

var (
	EmailExist      = &ServiceError{http.StatusConflict, "email already registered"}
	DuplicateEntry  = &ServiceError{http.StatusConflict, "data already exist"}
	UserNotActive   = &ServiceError{http.StatusForbidden, "user is not active"}
	InvalidPassword = &ServiceError{http.StatusUnauthorized, "invalid email or password"}
	TokenRequired   = &ServiceError{http.StatusUnauthorized, "token required"}
	TokenExpired    = &ServiceError{http.StatusUnauthorized, "token expired"}
	TokenInvalid    = &ServiceError{http.StatusUnauthorized, "token invalid"}
)
