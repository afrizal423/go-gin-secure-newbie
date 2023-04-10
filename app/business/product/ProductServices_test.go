package product_test

import (
	"errors"
	"testing"

	productServices "github.com/afrizal423/go-gin-secure-newbie/app/business/product"
	"github.com/afrizal423/go-gin-secure-newbie/app/models"
	productRepo "github.com/afrizal423/go-gin-secure-newbie/app/repository/product"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &productRepo.ProductRepositoryMock{Mock: mock.Mock{}}
var serviceProduct = productServices.ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	expectedProduct := models.Product{}
	var id uint = 90
	productRepository.Mock.On("GetDataById", id).Return(expectedProduct, errors.New("error"))

	product, err := serviceProduct.GetDataById(id)

	// Check if the mock repository was called
	productRepository.Mock.AssertCalled(t, "GetDataById", id)

	// Check if the returned products are correct
	assert.Error(t, err)
	assert.Equal(t, expectedProduct, product)
}

// TEST GET ONE PRODUCT FOUND
func TestProductServiceGetOneProduct(t *testing.T) {
	expectedProduct := models.Product{
		Title:       "Buku A",
		Description: "ini sebuah buku",
		UserID:      2,
	}
	var id uint = 2
	// Set up mock repository
	productRepository.Mock.On("GetDataById", id).Return(expectedProduct, nil)

	// Call service function
	product, err := serviceProduct.GetDataById(id)

	// Check if the mock repository was called
	productRepository.Mock.AssertCalled(t, "GetDataById", id)

	// Check if the returned products are correct
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	// expected result product
	expectedProduct := []models.Product{}
	productRepository.Mock.On("GetAll").Return(expectedProduct, errors.New("error"))

	products, err := serviceProduct.GetAll()

	assert.Error(t, err)
	assert.Len(t, products, 0)
	assert.Equal(t, expectedProduct, products)
}
func TestProductServiceGetAllProductFound(t *testing.T) {
	products := []models.Product{
		{
			Title:       "Buku A",
			Description: "ini sebuah buku",
		},
		{
			Title:       "Buku A",
			Description: "ini sebuah buku",
		},
	}

	productRepository.Mock.On("GetAll").Return(products, nil)

	result, err := serviceProduct.GetAll()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(products), len(result))
}
