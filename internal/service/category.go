package service

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/service_errors"
	"github.com/google/uuid"
)

type CategoryService struct {
	Repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		Repo: repo,
	}
}

func (s *CategoryService) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.CategoryResponsePagination, error) {
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

func (s *CategoryService) FindById(ctx context.Context, id uuid.UUID) (dto.CategoryResponse, error) {

	category, err := s.Repo.FindById(ctx, id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.ToCategoryResponse(*category), nil
}

func (s *CategoryService) Store(ctx context.Context, req dto.CreateCategoryRequest) (dto.CategoryResponse, error) {
	category := dto.ToCreateCategoryModel(req)

	if _, err := s.Repo.Store(ctx, &category); err != nil {
		if service_errors.IsUniqueViolation(err) {
			return dto.CategoryResponse{}, service_errors.DuplicateEntry
		}
		return dto.CategoryResponse{}, err
	}

	return dto.ToCategoryResponse(category), nil
}

func (s *CategoryService) Update(ctx context.Context, id uuid.UUID, req dto.UpdateCategoryRequest) (dto.CategoryResponse, error) {
	category := dto.ToUpdateCategoryModel(req)
	updated, err := s.Repo.Update(ctx, id, &category)
	if err != nil {
		if service_errors.IsUniqueViolation(err) {
			return dto.CategoryResponse{}, service_errors.DuplicateEntry
		}
		return dto.CategoryResponse{}, err
	}

	return dto.ToCategoryResponse(*updated), nil
}

func (s *CategoryService) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (dto.CategoryResponse, error) {
	category, err := s.Repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.ToCategoryResponse(*category), nil
}
