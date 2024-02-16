package repository

import "Hactiv8_Bootcamp/H9-mock/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
	FindAll() []*entity.Product
}
