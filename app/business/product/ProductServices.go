package product

import (
	"log"

	"github.com/afrizal423/go-gin-secure-newbie/app/models"
)

type ProductService struct {
	Repository IProductRepository
}

func NewProductService(repository IProductRepository) *ProductService {
	return &ProductService{
		repository,
	}
}

func (u *ProductService) GetAll() ([]models.Product, error) {
	if products, err := u.Repository.GetAll(); err != nil {
		log.Println("Data not found")
		return products, err
	} else {
		return products, nil
	}
}

func (u *ProductService) GetAllByUserId(userID uint) ([]models.Product, error) {
	if products, err := u.Repository.GetAllByUserId(userID); err != nil {
		log.Println("Data not found")
		return products, err
	} else {
		return products, nil
	}
}

func (u *ProductService) GetDataById(id uint) (models.Product, error) {
	if products, err := u.Repository.GetDataById(id); err != nil {
		log.Println("failed to create product")
		return products, err
	} else {
		return products, nil
	}
}

func (u *ProductService) GetDataByUserId(userID uint, id uint) (models.Product, error) {
	if products, err := u.Repository.GetDataByUserId(userID, id); err != nil {
		log.Println("failed to create product")
		return products, err
	} else {
		return products, nil
	}
}

func (u *ProductService) CreateProduct(product models.Product, userID uint) (models.Product, error) {
	if products, err := u.Repository.CreateProduct(product, userID); err != nil {
		log.Println("failed to create product")
		return products, err
	} else {
		return products, nil
	}
}

func (u *ProductService) DeleteProduct(id uint) (models.Product, error) {
	if products, err := u.Repository.DeleteProduct(id); err != nil {
		log.Println("failed to delete product")
		return products, err
	} else {
		return products, nil
	}
}

func (u *ProductService) UpdateProduct(product models.Product, id uint) (models.Product, error) {
	if product, err := u.Repository.UpdateProduct(product, id); err != nil {
		log.Println("Failed to update product")
		return product, err
	} else {
		return product, nil
	}
}
