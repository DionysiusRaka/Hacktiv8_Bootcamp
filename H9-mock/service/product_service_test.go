package service

import (
	"Hactiv8_Bootcamp/H9-mock/entity"
	"Hactiv8_Bootcamp/H9-mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}

var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "Product not found", err.Error(), "Error response has to be 'product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product{
		Id:   "2",
		Name: "Kacamata",
	}

	productRepository.Mock.On("FindById", "2").Return(product)
	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "Result has to be '2'")
	assert.Equal(t, product.Name, result.Name, "Result has to be 'Kacamata'")
	assert.Equal(t, &product, result, "Result has to be a product data with id '2'")
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindAll").Return(nil)

	product, err := productService.GetAllProduct()

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetAllProduct(t *testing.T) {
	product := []*entity.Product{
		{
			Id:   "1",
			Name: "Laptop",
		},
		{
			Id:   "2",
			Name: "Kacamata",
		},
	}

	productRepository.Mock.On("FindAll").Return(product)
	result, err := productService.GetAllProduct()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product, result, "result has to be products")
}
