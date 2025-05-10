package service

import (
	"category/internal/core/domain"
	"category/internal/core/port"
	"context"
	"fmt"
	"strconv"
)

type ProductService struct {
	client port.ProductClient
}

func NewProductService(client port.ProductClient) *ProductService {
	return &ProductService{
		client,
	}
}

func (s *ProductService) ListProducts(ctx context.Context, params domain.ListParamsSt, ids, categoryIDs []int64, withCategory bool) (*domain.ProductListRep, error) {
	stringIds := make([]string, len(ids))
	for i, id := range ids {
		stringIds[i] = strconv.FormatInt(id, 10)
	}

	stringCategoryIds := make([]string, len(categoryIDs))
	for i, id := range ids {
		stringIds[i] = strconv.FormatInt(id, 10)
	}

	return s.client.ListProducts(context.Background(), params, stringIds, stringCategoryIds, withCategory)
}

func (s *ProductService) GetProduct(ctx context.Context, id int64) (*domain.ProductMain, error) {
	if id < 0 {
		return nil, fmt.Errorf("id cannot be empty")
	}
	stringId := strconv.FormatInt(id, 10)

	return s.client.GetProduct(context.Background(), stringId)
}
