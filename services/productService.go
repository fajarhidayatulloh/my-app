package services

import (
	"gitlab.com/my-app/infrastructures"

	"gitlab.com/my-app/models"
	"gitlab.com/my-app/repositories"
)

type IProductService interface {
	StoreProduct(models.ProductInput) (models.Product, error)
	ProductList() ([]models.Product, error)
	GetProductByID(int) (models.Product, error)
}

type ProductService struct {
	ProductRepository repositories.IProductRepository
}

func InitProductService() *ProductService {
	productRepository := new(repositories.ProductRepository)
	productRepository.DB = &infrastructures.SQLConnection{}

	productService := new(ProductService)
	productService.ProductRepository = productRepository

	return productService
}

func (r *ProductService) StoreProduct(data models.ProductInput) (result models.Product, err error) {
	var product models.Product

	product.Name = data.Name
	product.Sku = data.Sku
	product.Price = data.Price
	product.Discount = data.Discount
	product.RewardPoint = data.RewardPoint
	product.RedeemPoint = data.RedeemPoint
	product.RedeemStamp = data.RedeemStamp
	product.Status = data.Status
	product.MerchantID = data.MerchantID
	product.CreatedAt = data.CreatedAt

	result, err = r.ProductRepository.StoreProduct(product)

	return
}

func (r *ProductService) ProductList() (result []models.Product, err error) {
	result, err = r.ProductRepository.ProductList()
	return
}

func (r *ProductService) GetProductByID(ID int) (result models.Product, err error) {
	result, err = r.ProductRepository.GetProductByID(ID)
	return
}
