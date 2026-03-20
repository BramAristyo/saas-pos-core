package usecase_errors

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
)

type UseCaseError struct {
	Code    int
	Message string
}

func (e *UseCaseError) Error() string {
	return e.Message
}

func IsUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

var (
	EmailExist      = &UseCaseError{http.StatusConflict, "email already registered"}
	DuplicateEntry  = &UseCaseError{http.StatusConflict, "data already exist"}
	UserNotActive   = &UseCaseError{http.StatusForbidden, "user is not active"}
	InvalidPassword = &UseCaseError{http.StatusUnauthorized, "invalid email or password"}
	TokenRequired   = &UseCaseError{http.StatusUnauthorized, "token required"}
	TokenExpired    = &UseCaseError{http.StatusUnauthorized, "token expired"}
	TokenInvalid    = &UseCaseError{http.StatusUnauthorized, "token invalid"}
	InvalidID       = &UseCaseError{http.StatusBadRequest, "invalid ID format"}
	NotFound        = &UseCaseError{http.StatusNotFound, "resource not found"}
)
