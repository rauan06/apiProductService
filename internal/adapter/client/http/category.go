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
	pathCategory = "category"
)

type CategoryClient struct {
	URL *url.URL
}

func NewCategoryClient(URL *url.URL) *CategoryClient {
	return &CategoryClient{
		URL.JoinPath(pathCategory),
	}
}

func (r *CategoryClient) ListCategories(ctx context.Context, params domain.ListParamsSt, ids []string) (*domain.CategoryListRep, error) {
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

	var categories = &domain.CategoryListRep{}
	if err := json.NewDecoder(resp.Body).Decode(categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryClient) GetCategory(ctx context.Context, id string) (*domain.CategoryMain, error) {
	resp, err := http.Get(r.URL.JoinPath(id).String())
	if err != nil {
		return nil, err
	}

	var categories *domain.CategoryMain
	if err := json.NewDecoder(resp.Body).Decode(categories); err != nil {
		return nil, err
	}

	return categories, nil
}
