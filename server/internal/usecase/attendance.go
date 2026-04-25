package usecase

import (
	"context"

	"github.com/BramAristyo/saas-pos-core/server/internal/api/dto"
	"github.com/BramAristyo/saas-pos-core/server/internal/repository"
	"github.com/BramAristyo/saas-pos-core/server/pkg/filter"
)

type AttendanceUseCase struct {
	Repo *repository.AttendanceRepository
}

func NewAttendanceUseCase(repo *repository.AttendanceRepository) *AttendanceUseCase {
	return &AttendanceUseCase{
		Repo: repo,
	}
}

func (u *AttendanceUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.AttendanceResponsePagination, error) {
	totalRows, attendances, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.AttendanceResponsePagination{}, err
	}

	res := dto.ToAttendanceResponses(attendances)
	return dto.ToAttendanceResponsePagination(res, req, totalRows), nil
}
