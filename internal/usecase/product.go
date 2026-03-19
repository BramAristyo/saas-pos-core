package usecase

import (
	"context"

	"github.com/BramAristyo/go-pos-mawish/internal/api/dto"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/pkg/filter"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
	"github.com/google/uuid"
)

type ProductUseCase struct {
	Repo *repository.ProductRepository
}

func NewProductUseCase(repo *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		Repo: repo,
	}
}

func (u *ProductUseCase) Paginate(ctx context.Context, req filter.PaginationWithInputFilter) (dto.ProductResponsePagination, error) {
	totalRows, products, err := u.Repo.Paginate(ctx, req)
	if err != nil {
		return dto.ProductResponsePagination{}, err
	}

	productResponses := make([]dto.ProductResponse, 0, len(products))
	for _, p := range products {
		productResponses = append(productResponses, dto.ToProductResponse(p))
	}

	return dto.ToProductResponsePagination(productResponses, req, totalRows), nil
}

func (u *ProductUseCase) FindById(ctx context.Context, id uuid.UUID) (dto.ProductResponse, error) {
	product, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ToProductResponse(product), nil
}

func (u *ProductUseCase) Store(ctx context.Context, req dto.CreateProductRequest) (dto.ProductResponse, error) {
	product := dto.ToProductModel(req)

	if _, err := u.Repo.Store(ctx, &product); err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.ProductResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.ProductResponse{}, err
	}

	return dto.ToProductResponse(product), nil
}

func (u *ProductUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateProductRequest) (dto.ProductResponse, error) {
	product := dto.ToUpdateProductModel(req)
	updated, err := u.Repo.Update(ctx, id, &product)
	if err != nil {
		if usecase_errors.IsUniqueViolation(err) {
			return dto.ProductResponse{}, usecase_errors.DuplicateEntry
		}
		return dto.ProductResponse{}, err
	}

	return dto.ToProductResponse(updated), nil
}

func (u *ProductUseCase) UpdateStatus(ctx context.Context, id uuid.UUID, status bool) (dto.ProductResponse, error) {
	product, err := u.Repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ToProductResponse(product), nil
}
