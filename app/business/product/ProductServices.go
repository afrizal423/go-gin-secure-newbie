package product

type ProductService struct {
	repository IProductRepository
}

func NewProductService(repository IProductRepository) *ProductService {
	return &ProductService{
		repository,
	}
}
