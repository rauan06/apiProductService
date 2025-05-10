package pipeline

import "category/internal/core/port"

type Pipeline struct {
	categorySvc port.CategoryService
	citySvc     port.CityService
	productSvc  port.ProductService
}

func NewPipeline(categorySvc port.CategoryService, citySvc port.CityService, productSvc port.ProductService) *Pipeline {
	return &Pipeline{
		categorySvc,
		citySvc,
		productSvc,
	}
}

func (p *Pipeline) Run() error { return nil }

func (p *Pipeline) Stop() {}
