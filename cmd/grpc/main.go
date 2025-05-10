package main

import (
	"context"
	"fmt"
	"log"

	"category/internal/adapter/client/grpc"
	"category/internal/core/domain"
	"category/internal/core/service"
)

const (
	apiURL = "api.com"
)

func main() {
	ctx := context.Background()

	// Initialize gRPC client
	client, err := grpc.NewCategoryClient(ctx, apiURL)
	if err != nil {
		log.Fatalf("Error initializing gRPC client: %v", err)
	}
	defer client.Close()

	// Initialize service with the client
	categoryService := service.NewCategoryService(client)

	// Extract data
	params := domain.ListParamsSt{
		Page:     1,
		PageSize: 10,
		Sort:     []string{"name"},
	}
	ids := []int64{123, 456}

	categories, err := categoryService.ListCategories(ctx, params, ids)
	if err != nil {
		log.Fatalf("Error extracting categories: %v", err)
	}

	// Load data into PostgreSQL (you need to implement the Postgres insertion logic here)

	fmt.Println("Successfully extracted and processed categories:", categories)
}
