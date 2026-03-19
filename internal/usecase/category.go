package usecase

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
	"github.com/google/uuid"
)

type CategoryUseCase struct {
	Repo *repository.CategoryRepository
}

func NewCategoryUseCase(repo *repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		Repo: repo,
	}
}

func (s *CategoryUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.CategoryResponsePagination, error) {
	totalRows, categories, err := s.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.CategoryResponsePagination{}, err
	}

	categoriesResponses := make([]dto.CategoryResponse, 0, len(categories))
	for _, c := range categories {
		categoriesResponses = append(categoriesResponses, dto.ToCategoryResponse(c))
	}

	return dto.ToCategoryResponsePagination(categoriesResponses, req, totalRows), nil
}

func (s *CategoryUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.CategoryResponse, error) {

	category, err := s.Repo.FindById(ctx, id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.ToCategoryResponse(*category), nil
}

func (s *CategoryUseCase) Store(ctx context.Context, req dto.CreateCategoryRequest) (dto.CategoryResponse, error) {
	category := dto.ToCreateCategoryModel(req)

	if _, err := s.Repo.Store(ctx, &category); err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.CategoryResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.CategoryResponse{}, err
	}

	return dto.ToCategoryResponse(category), nil
}

func (s *CategoryUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateCategoryRequest) (dto.CategoryResponse, error) {
	category := dto.ToUpdateCategoryModel(req)
	updated, err := s.Repo.Update(ctx, id, &category)
	if err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.CategoryResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.CategoryResponse{}, err
	}

	return dto.ToCategoryResponse(*updated), nil
}

func (s *CategoryUseCase) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (dto.CategoryResponse, error) {
	category, err := s.Repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.ToCategoryResponse(*category), nil
}
