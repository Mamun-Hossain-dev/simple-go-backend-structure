package product

import "fmt"

type ProductService interface {
	GetProducts() []Product
	GetProductByID(id int) (*Product, error)
	CreateProduct(p Product) Product
	UpdateProduct(id int, p Product) (*Product, error)
	DeleteProduct(id int) (*Product, error)
}

type productService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProducts() []Product {
	return s.repo.GetAll()
}

func (s *productService) GetProductByID(id int) (*Product, error) {
	p := s.repo.GetByID(id)
	if p == nil {
		return nil, fmt.Errorf("product not found")
	}
	return p, nil
}

func (s *productService) CreateProduct(p Product) Product {
	return s.repo.Store(p)
}

func (s *productService) UpdateProduct(id int, p Product) (*Product, error) {
	updatedProduct := s.repo.Update(id, p)
	if updatedProduct == nil {
		return nil, fmt.Errorf("product not found")
	}
	return updatedProduct, nil
}

func (s *productService) DeleteProduct(id int) (*Product, error) {
	deletedProduct := s.repo.Delete(id)
	if deletedProduct == nil {
		return nil, fmt.Errorf("product not found")
	}
	return deletedProduct, nil
}
