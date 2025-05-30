package service

import (
	"context"
	"encoding/json"
	"github.com/Deevins/lampshop-backend/product/internal/model"
	"github.com/Deevins/lampshop-backend/product/internal/service/sql"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"time"
)

type ProductService struct {
	db *pgxpool.Pool
}

func NewProductService(db *pgxpool.Pool) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) ListProducts(ctx context.Context) ([]*sql.Product, error) {
	repo := sql.New(s.db)
	return repo.ListProducts(ctx)
}

func (s *ProductService) GetProduct(ctx context.Context, id uuid.UUID) (*sql.Product, error) {
	repo := sql.New(s.db)
	product, err := repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, params *model.CreateProductRequest) error {
	repo := sql.New(s.db)

	raw, err := json.Marshal(params.Attributes)
	if err != nil {
		return err
	}

	return repo.CreateProduct(ctx, &sql.CreateProductParams{
		ID:          uuid.New(),
		Sku:         params.SKU,
		Name:        params.Name,
		Description: lo.ToPtr(params.Description),
		CategoryID:  params.CategoryID,
		IsActive:    params.IsActive,
		ImageUrl:    lo.ToPtr(params.ImageURL),
		Price:       decimal.NewFromFloat(params.Price),
		StockQty:    int32(params.StockQty),
		Attributes:  raw,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
}

func (s *ProductService) UpdateProduct(ctx context.Context, params *model.UpdateProductRequest) error {
	repo := sql.New(s.db)
	return repo.UpdateProduct(ctx, mapUpdateRequestToSQL(params.ID, params))
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	repo := sql.New(s.db)
	return repo.DeleteProduct(ctx, id)
}

func (s *ProductService) ListCategories(ctx context.Context) ([]*sql.Category, error) {
	repo := sql.New(s.db)
	return repo.ListCategories(ctx)
}

func (s *ProductService) CreateCategory(ctx context.Context, c *model.CategoryCreateRequest) error {
	repo := sql.New(s.db)
	return repo.CreateCategory(ctx, &sql.CreateCategoryParams{
		ID:   uuid.New(),
		Name: c.Name,
	})
}

func (s *ProductService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	repo := sql.New(s.db)
	return repo.DeleteCategory(ctx, id)
}
