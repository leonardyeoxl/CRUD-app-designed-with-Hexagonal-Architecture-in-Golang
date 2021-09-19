package tests

import (
	"testing"

	"github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/domain/mocks"
    "github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/domain"
)

type SectionTester struct {
    sectionService *mocks.SectionService
}

func (tester *SectionTester) Create(name string, code string) {
    tester.sectionService.Store(name, code)
}

func (tester *SectionTester) Read(code string) {
    tester.sectionService.Find(code)
}

func (tester *SectionTester) Update(code string, name string) {
    tester.sectionService.Update(code, name)
}

func (tester *SectionTester) FindAll() {
    tester.sectionService.FindAll()
}

func (tester *SectionTester) Delete(code string) {
    tester.sectionService.Delete(code)
}

func TestCreateSection(t *testing.T) {
    mockSectionService := &mocks.SectionService{}
    testSectionService := &SectionTester{sectionService:mockSectionService}
    mockSectionService.On("Store", "abc", "04cff685-5fbc-4cb9-a08e-a9154fc58cd5").Return(nil).Once()
    testSectionService.Create("abc", "04cff685-5fbc-4cb9-a08e-a9154fc58cd5")
    mockSectionService.AssertExpectations(t)
}

func TestReadSection(t *testing.T) {
    mockSectionService := &mocks.SectionService{}
    testSectionService := &SectionTester{sectionService:mockSectionService}
    mockSectionService.On("Find", "04cff685-5fbc-4cb9-a08e-a9154fc58cd5").Return(domain.Section{
        Code: "04cff685-5fbc-4cb9-a08e-a9154fc58cd5",
        Name: "abc",
    }, nil).Once()
    testSectionService.Read("04cff685-5fbc-4cb9-a08e-a9154fc58cd5")
    mockSectionService.AssertExpectations(t)
}

func TestUpdateSection(t *testing.T) {
    mockSectionService := &mocks.SectionService{}
    testSectionService := &SectionTester{sectionService:mockSectionService}
    mockSectionService.On("Update", "04cff685-5fbc-4cb9-a08e-a9154fc58cd5", "def").Return(nil).Once()
    testSectionService.Update("04cff685-5fbc-4cb9-a08e-a9154fc58cd5", "def")
    mockSectionService.AssertExpectations(t)
}

func TestFindAllSections(t *testing.T) {
    merchant_01 := domain.Merchant{
        Code: "04cff685-5fbc-4cb9-a08e-a9154fc58cd5",
        Name: "def",
    }
    mockSectionService := &mocks.SectionService{}
    testSectionService := &SectionTester{sectionService:mockSectionService}
    mockSectionService.On("FindAll").Return([]domain.Merchant{merchant_01}, nil).Once()
    testSectionService.FindAll()
    mockSectionService.AssertExpectations(t)
}

func TestDeleteSection(t *testing.T) {
    mockSectionService := &mocks.SectionService{}
    testSectionService := &SectionTester{sectionService:mockSectionService}
    mockSectionService.On("Delete", "04cff685-5fbc-4cb9-a08e-a9154fc58cd5").Return(nil).Once()
    testSectionService.Delete("04cff685-5fbc-4cb9-a08e-a9154fc58cd5")
    mockSectionService.AssertExpectations(t)
}
