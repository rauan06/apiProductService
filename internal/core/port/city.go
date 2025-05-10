package port

import (
	"context"

	"category/internal/core/domain"
)

type CityService interface {
	ListCities(ctx context.Context, params domain.ListParamsSt, ids []int64) (*domain.CityListRep, error)
	GetCity(ctx context.Context, id int64) (*domain.CityMain, error)
}

type CityClient interface {
	ListCities(ctx context.Context, params domain.ListParamsSt, ids []string) (*domain.CityListRep, error)
	GetCity(ctx context.Context, id string) (*domain.CityMain, error)
}
