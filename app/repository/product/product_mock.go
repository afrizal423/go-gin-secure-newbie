package product

import (
	"github.com/afrizal423/go-gin-secure-newbie/app/models"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *models.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(models.Product)

	return &product
}

func (repository *ProductRepositoryMock) GetAll() ([]models.Product, error) {
	args := repository.Mock.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (repository *ProductRepositoryMock) GetAllByUserId(userID uint) ([]models.Product, error) {
	args := repository.Mock.Called(userID)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (repository *ProductRepositoryMock) GetDataByUserId(userID uint, id uint) (models.Product, error) {
	args := repository.Mock.Called(userID, id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (repository *ProductRepositoryMock) GetDataById(id uint) (models.Product, error) {
	args := repository.Mock.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (repository *ProductRepositoryMock) CreateProduct(product models.Product, userID uint) (models.Product, error) {
	args := repository.Mock.Called(product, userID)
	return args.Get(0).(models.Product), args.Error(1)
}

func (repository *ProductRepositoryMock) UpdateProduct(product models.Product, id uint) (models.Product, error) {
	args := repository.Mock.Called(product, id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (repository *ProductRepositoryMock) DeleteProduct(id uint) (models.Product, error) {
	args := repository.Mock.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}
