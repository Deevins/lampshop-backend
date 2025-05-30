package service

import (
	"context"
	"github.com/Deevins/lampshop-backend/product/internal/service/sql"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (s *ProductService) CreateProduct(ctx context.Context, params *sql.CreateProductParams) error {
	repo := sql.New(s.db)
	return repo.CreateProduct(ctx, params)
}

func (s *ProductService) UpdateProduct(ctx context.Context, params *sql.UpdateProductParams) error {
	repo := sql.New(s.db)
	return repo.UpdateProduct(ctx, params)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	repo := sql.New(s.db)
	return repo.DeleteProduct(ctx, id)
}

func (s *ProductService) ListCategories(ctx context.Context) ([]*sql.Category, error) {
	repo := sql.New(s.db)
	return repo.ListCategories(ctx)
}

func (s *ProductService) CreateCategory(ctx context.Context, c *sql.CreateCategoryParams) error {
	repo := sql.New(s.db)
	return repo.CreateCategory(ctx, c)
}

func (s *ProductService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	repo := sql.New(s.db)
	return repo.DeleteCategory(ctx, id)
}
