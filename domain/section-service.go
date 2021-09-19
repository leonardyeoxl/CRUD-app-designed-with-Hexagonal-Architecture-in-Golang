package domain

type SectionService interface {
	Find(code string) (Section, error)
	Store(name string, code string) error
	Update(code string, name string) error
	FindAll() ([]Section, error)
	Delete(code string) error
}

type SectionRepository interface {
	FindSection(code string) (Section, error)
	StoreSection(name string, code string) error
	UpdateSection(code string, name string) error
	FindAllSection() ([]Section, error)
	DeleteSection(code string) error
}

type section_service struct {
	sectionrepo SectionRepository
}

func NewSectionService(sectionrepo SectionRepository) SectionService {
	return &section_service{sectionrepo: sectionrepo}
}

func (ss *section_service) Find(code string) (Section, error) {
	return ss.section_service.FindSection(code)
}

func (ss *section_service) Store(name string, code string) error {
	return ss.section_service.StoreSection(name, code)
}

func (ss *section_service) Update(code string, name string) error {
	return ss.section_service.UpdateSection(code, name) 
}

func (ss *section_service) FindAll() ([]Section, error) {
	return ss.section_service.FindAllSection() 
}

func (ss *section_service) Delete(code string) error {
	return ss.section_service.DeleteSection(code)
}