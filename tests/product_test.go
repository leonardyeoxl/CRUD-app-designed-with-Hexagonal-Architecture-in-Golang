package tests

import (
	"testing"

	"github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/domain/mocks"
    "github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/domain"
)

type ProductTester struct {
    productService *mocks.ProductService
}

func (tester *ProductTester) Create(email string, name string, merchant_code string) {
    tester.productService.Store(email, name, merchant_code)
}

func (tester *ProductTester) Read(email string) {
    tester.productService.Find(email)
}

func (tester *ProductTester) Update(email string, name string) {
    tester.productService.Update(email, name)
}

func (tester *ProductTester) FindAll(merchant_code string, page int, page_size int) {
    tester.productService.FindAll(merchant_code, page, page_size)
}

func (tester *ProductTester) Delete(email string) {
    tester.productService.Delete(email)
}

func TestCreateProduct(t *testing.T) {
    mockProductService := &mocks.ProductService{}
    testProductService := &ProductTester{productService:mockProductService}
    mockProductService.On("Store", "jim@abc.com", "jim", "04cff685-5fbc-4cb9-a08e-a9154fc58cd5").Return(nil).Once()
    testProductService.Create("jim@abc.com", "jim", "04cff685-5fbc-4cb9-a08e-a9154fc58cd5")
    mockProductService.AssertExpectations(t)
}

func TestReadProduct(t *testing.T) {
    mockProductService := &mocks.ProductService{}
    testProductService := &ProductTester{productService:mockProductService}
    mockProductService.On("Find", "jim@abc.com").Return(domain.TeamMember{
        Email: "jim@abc.com",
        Name: "jim",
    }, nil).Once()
    testProductService.Read("jim@abc.com")
    mockProductService.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
    mockProductService := &mocks.ProductService{}
    testProductService := &ProductTester{productService:mockProductService}
    mockProductService.On("Update", "jim@abc.com", "tim").Return(nil).Once()
    testProductService.Update("jim@abc.com", "tim")
    mockProductService.AssertExpectations(t)
}

func TestFindAllProducts(t *testing.T) {
    mockProductService := &mocks.ProductService{}
    testProductService := &ProductTester{productService:mockProductService}
    mockProductService.On("FindAll", "04cff685-5fbc-4cb9-a08e-a9154fc58cd5", 1, 2).Return([]domain.TeamMember{}, nil).Once()
    testProductService.FindAll("04cff685-5fbc-4cb9-a08e-a9154fc58cd5", 1, 2)
    mockProductService.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
    mockProductService := &mocks.ProductService{}
    testProductService := &ProductTester{productService:mockProductService}
    mockProductService.On("Delete", "jim@abc.com").Return(nil).Once()
    testProductService.Delete("jim@abc.com")
    mockProductService.AssertExpectations(t)
}
