package product

import (
	"log"

	"github.com/afrizal423/go-gin-secure-newbie/app/models"
)

type ProductService struct {
	repository IProductRepository
}

func NewProductService(repository IProductRepository) *ProductService {
	return &ProductService{
		repository,
	}
}

func (u *ProductService) GetAll() ([]models.Product, error) {
	if products, err := u.repository.GetAll(); err != nil {
		log.Println("Data not found")
		return products, err
	} else {
		return products, nil
	}
}

func (u *ProductService) GetAllByUserId(userID uint) ([]models.Product, error) {
	if products, err := u.repository.GetAllByUserId(userID); err != nil {
		log.Println("Data not found")
		return products, err
	} else {
		return products, nil
	}
}

func (u *ProductService) GetDataById(id uint) (models.Product, error) {
	if products, err := u.repository.GetDataById(id); err != nil {
		log.Println("failed to create product")
		return products, err
	} else {
		return products, nil
	}
}

func (u *ProductService) GetDataByUserId(userID uint, id uint) (models.Product, error) {
	if products, err := u.repository.GetDataByUserId(userID, id); err != nil {
		log.Println("failed to create product")
		return products, err
	} else {
		return products, nil
	}
}

func (u *ProductService) CreateProduct(product models.Product, userID uint) (models.Product, error) {
	if products, err := u.repository.CreateProduct(product, userID); err != nil {
		log.Println("failed to create product")
		return products, err
	} else {
		return products, nil
	}
}
