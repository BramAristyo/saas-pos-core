package helper

import (
	"context"

	"github.com/BramAristyo/saas-pos-core/server/internal/constant"
	"github.com/BramAristyo/saas-pos-core/server/pkg/usecase_errors"
	"github.com/google/uuid"
)

func ExtractUserID(ctx context.Context) (uuid.UUID, error) {
	val := ctx.Value(constant.CtxUserID)
	if val == nil {
		return uuid.Nil, usecase_errors.TokenRequired
	}

	strVal, ok := val.(string)
	if !ok {
		return uuid.Nil, usecase_errors.TokenInvalid
	}

	id, err := uuid.Parse(strVal)
	if err != nil {
		return uuid.Nil, usecase_errors.InvalidID
	}

	return id, nil
}
