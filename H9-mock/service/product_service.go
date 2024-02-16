package service

import (
	"Hactiv8_Bootcamp/H9-mock/entity"
	"Hactiv8_Bootcamp/H9-mock/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("Product not found")
	}

	return product, nil
}

func (service ProductService) GetAllProduct() ([]*entity.Product, error) {
	product := service.Repository.FindAll()
	if len(product) != 2 {
		return nil, errors.New("Product not found")
	}

	return product, nil
}
