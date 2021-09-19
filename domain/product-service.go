package domain

type ProductService interface {
	Find(code string) (Product, error)
	Store(code string, name string, price float64, quantity int, section_code string) error
	Update(code string, name string, price float64, quantity int) error
	FindAll(merchant_code string, page int, page_size int) ([]Product, error)
	Delete(code string) error
}

type ProductRepository interface {
	FindProduct(code string) (Product, error)
	StoreProduct(code string, name string, price float64, quantity int, section_code string) error
	UpdateProduct(code string, name string, price float64, quantity int) error
	FindAllProduct(section_code string, page int, page_size int) ([]Product, error)
	DeleteProduct(code string) error
}

type product_service struct {
	productrepo ProductRepository
}

func NewProductService(productrepo ProductRepository) ProductService {
	return &product_service{productrepo: productrepo}
}

func (ps *product_service) Find(code string) (Product, error) {
	return ps.productrepo.FindProduct(code)
}

func (ps *product_service) Store(code string, name string, price float64, quantity int, section_code string) error {
	return ps.productrepo.StoreProduct(code, name, price, quantity, section_code)
}

func (ps *product_service) Update(code string, name string, price float64, quantity int) error {
	return ps.productrepo.UpdateProduct(code, name, price, quantity) 
}

func (ps *product_service) FindAll(section_code string, page int, page_size int) ([]Product, error) {
	return ps.productrepo.FindAllProduct(section_code, page, page_size) 
}

func (ps *product_service) Delete(code string) error {
	return ps.productrepo.DeleteProduct(code)
}