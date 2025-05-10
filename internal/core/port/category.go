package port

import (
	"context"

	"category/internal/core/domain"
)

type CategoryService interface {
	ListCategories(ctx context.Context, params domain.ListParamsSt, ids []int64) (*domain.CategoryListRep, error)
	GetCategory(ctx context.Context, id int64) (*domain.CategoryMain, error)
}

type CategoryClient interface {
	ListCategories(ctx context.Context, params domain.ListParamsSt, ids []string) (*domain.CategoryListRep, error)
	GetCategory(ctx context.Context, id string) (*domain.CategoryMain, error)
}
