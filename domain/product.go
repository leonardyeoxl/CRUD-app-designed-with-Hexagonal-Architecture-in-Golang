package domain

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID				int64		`json:"id"  gorm:"primary_key"`
	Name			string		`json:"name" gorm:"column:name"`
	Price			float64		`json:"price"`
	Quantity		int			`json:"quantity"`
    SectionID 		int 		`json:"sectionID"`
    Section 		Section 	`json:"section"`
}