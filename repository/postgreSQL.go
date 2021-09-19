package repository

import (
	"github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/domain"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func newPostgreSQLDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	// Migrate the schema
	db.AutoMigrate(&domain.Section{}, &domain.Product{})

	return db, nil
}

func NewPostgreSQLRepository(dsn string) (domain.SectionRepository, domain.ProductRepository, error) {
	postgreSQLDB, err := newPostgreSQLDB(dsn)
	repo := &postgreSQLRepository{
		db: postgreSQLDB,
	}
	if err != nil {
		return nil, nil, errors.Wrap(err, "DB error")
	}

	return repo, repo, nil
}

func Paginate(page int, page_size int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
	  	}
  
		offset := (page - 1) * page_size
		return db.Offset(offset).Limit(page_size)
	}
}

func (postgreSQLRepo *postgreSQLRepository) StoreSection(name string, uuid string) error {
	result := postgreSQLRepo.db.Create(&domain.Section{ 
		Code: uuid,
		Name: name,
	})
	if result.Error != nil {
		return result.Error
	}
	
	return nil
}

func (postgreSQLRepo *postgreSQLRepository) UpdateSection(code string, name string) error {
	section := domain.Section{}
	result := postgreSQLRepo.db.Model(&section).Where("code = ?", code).Update("name", name)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (postgreSQLRepo *postgreSQLRepository) FindSection(code string) (domain.Section, error) {
	section := domain.Section{}
	result := postgreSQLRepo.db.First(&section, "code = ?", code)
	if result.Error != nil {
		return section, result.Error
	}

	return section, nil
}

func (postgreSQLRepo *postgreSQLRepository) FindAllSection() ([]domain.Section, error) {
	sections := []domain.Section{}
	result := postgreSQLRepo.db.Find(&sections)
	if result.Error != nil {
		return sections, result.Error
	}

	return sections, nil
}

func (postgreSQLRepo *postgreSQLRepository) DeleteSection(code string) error {
	section := domain.Section{}
	result := postgreSQLRepo.db.Where("code = ?", code).Delete(&section)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (postgreSQLRepo *postgreSQLRepository) StoreProduct(code string, name string, price float64, quantity int, section_code string) error {
	section := domain.Section{}
	result := postgreSQLRepo.db.First(&section, "code = ?", section_code)
	if result.Error != nil {
		return result.Error
	}

	product := domain.Product{
		Name: name,
		Price: price,
		Quantity: quantity,
		Section: section,
	}
	product_result := postgreSQLRepo.db.Create(&product)
	if product_result.Error != nil {
		return product_result.Error
	}
	
	return nil
}

func (postgreSQLRepo *postgreSQLRepository) UpdateProduct(code string, name string, price float64, quantity int) error {
	product := domain.Product{}
	result := postgreSQLRepo.db.Model(&product).Where("code = ?", code).Omit("Section").Updates(domain.Product{
		Name: name,
		Price: price,
		Quantity: quantity,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (postgreSQLRepo *postgreSQLRepository) FindProduct(code string) (domain.Product, error) {
	product := domain.Product{}
	result := postgreSQLRepo.db.Preload("Section").First(&product, "code = ?", code)
	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}

func (postgreSQLRepo *postgreSQLRepository) FindAllProduct(section_code string, page int, page_size int) ([]domain.Product, error) {
	products := []domain.Product{}
	section := domain.Section{}
	result := postgreSQLRepo.db.First(&section, "code = ?", section_code)
	if result.Error != nil {
		return products, result.Error
	}

    postgreSQLRepo.db.Preload("Section", "code = ?", section_code).Scopes(Paginate(page, page_size)).Find(&products)
	return products, nil
}

func (postgreSQLRepo *postgreSQLRepository) DeleteProduct(code string) error {
	product := domain.Product{}
	result := postgreSQLRepo.db.Where("code = ?", code).Delete(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}