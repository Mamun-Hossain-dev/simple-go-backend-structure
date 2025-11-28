package product

import "fmt"

type ProductService interface {
	GetProducts() ([]Product, error)
	GetProductByID(id int) (*Product, error)
	CreateProduct(p Product) (*Product, error)
	UpdateProduct(id int, p Product) (*Product, error)
	DeleteProduct(id int) (*Product, error)
}

type productService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProducts() ([]Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}
	return products, nil
}

func (s *productService) GetProductByID(id int) (*Product, error) {
	p, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}
	if p == nil {
		return nil, fmt.Errorf("product not found")
	}
	return p, nil
}

func (s *productService) CreateProduct(p Product) (*Product, error) {
	createdProduct, err := s.repo.Store(p)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}
	return createdProduct, nil
}

func (s *productService) UpdateProduct(id int, p Product) (*Product, error) {
	updatedProduct, err := s.repo.Update(id, p)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}
	if updatedProduct == nil {
		return nil, fmt.Errorf("product not found")
	}
	return updatedProduct, nil
}

func (s *productService) DeleteProduct(id int) (*Product, error) {
	deletedProduct, err := s.repo.Delete(id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete product: %w", err)
	}
	if deletedProduct == nil {
		return nil, fmt.Errorf("product not found")
	}
	return deletedProduct, nil
}
