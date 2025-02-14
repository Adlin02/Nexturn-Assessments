package service

import (
	"ecommerce-inventory/model"
	"ecommerce-inventory/repository"
	"errors"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// AddProduct adds a new product to the inventory.
func (service *ProductService) AddProduct(product *model.Product) error {
	// Validate product data
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return errors.New("invalid product data")
	}

	// Insert product into the database
	return service.repo.AddProduct(product)
}

// GetProductByID retrieves a product by its ID.
func (service *ProductService) GetProductByID(id int) (*model.Product, error) {
	product, err := service.repo.GetProductByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

// UpdateProduct updates a product's details.
func (service *ProductService) UpdateProduct(product *model.Product) error {
	// Validate product data
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return errors.New("invalid product data")
	}

	// Update product in the database
	return service.repo.UpdateProduct(product)
}

// DeleteProduct deletes a product from the inventory.
func (service *ProductService) DeleteProduct(id int) error {
	return service.repo.DeleteProduct(id)
}

// GetAllProducts retrieves all products with pagination.
func (service *ProductService) GetAllProducts(page, limit int) ([]model.Product, error) {
	return service.repo.GetAllProducts(page, limit)
}
