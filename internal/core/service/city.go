package service

import (
	"category/internal/core/domain"
	"category/internal/core/port"
	"context"
	"fmt"
	"strconv"
)

type CityService struct {
	client port.CityClient
}

func NewCityService(client port.CityClient) *CityService {
	return &CityService{
		client,
	}
}

func (s *CityService) ListCities(ctx context.Context, params domain.ListParamsSt, ids []int64) (*domain.CityListRep, error) {
	stringIds := make([]string, len(ids))
	for i, id := range ids {
		stringIds[i] = strconv.FormatInt(id, 10)
	}

	return s.client.ListCities(context.Background(), params, stringIds)
}

func (s *CityService) GetCity(ctx context.Context, id int64) (*domain.CityMain, error) {
	if id < 0 {
		return nil, fmt.Errorf("id cannot be empty")
	}
	stringId := strconv.FormatInt(id, 10)

	return s.client.GetCity(context.Background(), stringId)
}
