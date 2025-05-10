package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"category/internal/core/domain"
)

const (
	pathCity = "city"
)

type CityClient struct {
	URL *url.URL
}

func NewCityClient(URL *url.URL) *CityClient {
	return &CityClient{
		URL.JoinPath(pathCity),
	}
}

func (r *CityClient) ListCities(ctx context.Context, params domain.ListParamsSt, ids []string) (*domain.CityListRep, error) {
	// Format params to url.Values format
	urlParams := r.URL.Query()
	urlParams.Set("list_params.page", strconv.FormatInt(params.Page, 10))
	urlParams.Set("list_params.page_size", strconv.FormatInt(params.PageSize, 10))

	for _, sortVal := range params.Sort {
		urlParams.Add("list_params.sort", sortVal)
	}

	for _, id := range ids {
		urlParams.Add("ids", id)
	}

	r.URL.RawQuery = urlParams.Encode()

	// Make a request
	resp, err := http.Get(r.URL.String())
	if err != nil {
		return nil, err
	}

	var cities = &domain.CityListRep{}
	if err := json.NewDecoder(resp.Body).Decode(cities); err != nil {
		return nil, err
	}

	return cities, nil
}

func (r *CityClient) GetCity(ctx context.Context, id string) (*domain.CityMain, error) {
	resp, err := http.Get(r.URL.JoinPath(id).String())
	if err != nil {
		return nil, err
	}

	var categories *domain.CityMain
	if err := json.NewDecoder(resp.Body).Decode(categories); err != nil {
		return nil, err
	}

	return categories, nil
}
