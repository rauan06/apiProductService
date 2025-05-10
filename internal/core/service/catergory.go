package service

import (
	"category/internal/core/domain"
	"category/internal/core/port"
	"context"
	"fmt"
	"strconv"
)

type CategoryService struct {
	client port.CategoryClient
}

func NewCategoryService(client port.CategoryClient) *CategoryService {
	return &CategoryService{
		client,
	}
}

func (s *CategoryService) ListCategories(ctx context.Context, params domain.ListParamsSt, ids []int64) (*domain.CategoryListRep, error) {
	stringIds := make([]string, len(ids))
	for i, id := range ids {
		stringIds[i] = strconv.FormatInt(id, 10)
	}

	return s.client.ListCategories(context.Background(), params, stringIds)
}

func (s *CategoryService) GetCategory(ctx context.Context, id int64) (*domain.CategoryMain, error) {
	if id < 0 {
		return nil, fmt.Errorf("id cannot be empty")
	}
	stringId := strconv.FormatInt(id, 10)

	return s.client.GetCategory(context.Background(), stringId)
}
