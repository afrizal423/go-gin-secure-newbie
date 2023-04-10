package product

import "github.com/afrizal423/go-gin-secure-newbie/app/models"

type IProductService interface {
	// get all data product
	GetAll() ([]models.Product, error)
	// get all data product by user id
	GetAllByUserId(userID uint) ([]models.Product, error)
	// create data product
	CreateProduct(product models.Product, userID uint) (models.Product, error)
	// get detail data by id (admin)
	GetDataById(id uint) (models.Product, error)
	// get detail data by id (user)
	GetDataByUserId(userID uint, id uint) (models.Product, error)
	// hapus data produk
	DeleteProduct(id uint) (models.Product, error)
	// update data produk
	UpdateProduct(product models.Product, id uint) (models.Product, error)
}

type IProductRepository interface {
	// get all data product
	GetAll() ([]models.Product, error)
	// get all data product by user id
	GetAllByUserId(userID uint) ([]models.Product, error)
	// create data product
	CreateProduct(product models.Product, userID uint) (models.Product, error)
	// get detail data by id (admin)
	GetDataById(id uint) (models.Product, error)
	// get detail data by id (user)
	GetDataByUserId(userID uint, id uint) (models.Product, error)
	// hapus data produk
	DeleteProduct(id uint) (models.Product, error)
	// update data produk
	UpdateProduct(product models.Product, id uint) (models.Product, error)
}
