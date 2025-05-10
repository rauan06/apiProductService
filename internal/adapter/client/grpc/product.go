package grpc

import (
	"context"

	"category/internal/core/domain"
	pb "category/protos/product/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	conn    *grpc.ClientConn
	service pb.ProductClient
}

func NewProductClient(ctx context.Context, target string) (*ProductClient, error) {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &ProductClient{
		conn:    conn,
		service: pb.NewProductClient(conn),
	}, nil
}

func (c *ProductClient) Close() {
	c.conn.Close()
}

func (c *ProductClient) ListProducts(ctx context.Context, params domain.ListParamsSt, ids, categoryIDs []string, withCategory bool) (*domain.ProductListRep, error) {
	resp, err := c.service.List(ctx, &pb.ProductListReq{
		ListParams: &pb.ListParamsSt{
			Page:     params.Page,
			PageSize: params.PageSize,
			Sort:     params.Sort,
		},
		Ids:          ids,
		CategoryIds:  categoryIDs,
		WithCategory: withCategory,
	})
	if err != nil {
		return nil, err
	}

	var results []domain.ProductMain
	for _, prod := range resp.Results {
		results = append(results, domain.ProductMain{
			UpdatedAt:   prod.UpdatedAt.String(),
			CreatedAt:   prod.CreatedAt.String(),
			Deleted:     prod.Deleted,
			ID:          prod.Id,
			Name:        prod.Name,
			Description: prod.Description,
			CategoryID:  prod.CategoryId,
			Category: domain.CategoryMain{
				UpdatedAt: prod.Category.UpdatedAt.String(),
				CreatedAt: prod.Category.CreatedAt.String(),
				Deleted:   prod.Category.Deleted,
				ID:        prod.Category.Id,
				Name:      prod.Category.Name,
			},
		})
	}

	return &domain.ProductListRep{
		PaginationInfo: domain.PaginationInfoSt{
			Page:     resp.PaginationInfo.Page,
			PageSize: resp.PaginationInfo.PageSize,
		},
		Results: results,
	}, nil
}

func (c *ProductClient) GetProduct(ctx context.Context, id string) (*domain.ProductMain, error) {
	resp, err := c.service.Get(ctx, &pb.ProductGetReq{Id: id})
	if err != nil {
		return nil, err
	}
	return &domain.ProductMain{
		CreatedAt:   resp.CreatedAt.String(),
		UpdatedAt:   resp.UpdatedAt.String(),
		Deleted:     resp.Deleted,
		ID:          resp.Id,
		Name:        resp.Name,
		Description: resp.Description,
		CategoryID:  resp.CategoryId,
		Category: domain.CategoryMain{
			CreatedAt: resp.Category.CreatedAt.String(),
			UpdatedAt: resp.Category.UpdatedAt.String(),
			Deleted:   resp.Category.Deleted,
			ID:        resp.Category.Id,
			Name:      resp.Category.Name,
		},
	}, nil
}
