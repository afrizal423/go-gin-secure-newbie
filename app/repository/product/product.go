package product

import (
	"log"

	"github.com/afrizal423/go-gin-secure-newbie/app/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db,
	}
}

func (u *ProductRepository) GetAll() ([]models.Product, error) {
	products := []models.Product{}
	if err := u.db.Order("id DESC").Find(&products).Error; err != nil {
		log.Println("Error get all products:", err)
		return products, err
	}
	return products, nil
}

func (u *ProductRepository) GetAllByUserId(userID uint) ([]models.Product, error) {
	products := []models.Product{}
	if err := u.db.Where("user_id = ?", userID).Order("id DESC").Find(&products).Error; err != nil {
		log.Println("Error get all products:", err, " by id:", userID)
		return products, err
	}
	return products, nil
}

func (u *ProductRepository) GetDataById(id uint) (models.Product, error) {
	product := models.Product{}
	if err := u.db.Where("id = ?", id).First(&product).Error; err != nil {
		log.Println("Error finding product:", err)
		return product, err
	}
	return product, nil
}

func (u *ProductRepository) GetDataByUserId(userID uint, id uint) (models.Product, error) {
	product := models.Product{}
	if err := u.db.Where("user_id = ? AND id = ?", userID, id).First(&product).Error; err != nil {
		log.Println("Error finding product:", err)
		return product, err
	}
	return product, nil
}

func (u *ProductRepository) CreateProduct(product models.Product, userID uint) (models.Product, error) {
	product.UserID = userID
	if err := u.db.Create(&product).Error; err != nil {
		log.Println("Error creating product:", err)
		return product, err
	}
	return product, nil
}
