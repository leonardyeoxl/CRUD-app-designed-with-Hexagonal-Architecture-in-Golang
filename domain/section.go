package domain

import (
	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	ID			int64
	Code		string			`json:"code" gorm:"column:code"`
	Name		string			`json:"name" gorm:"column:email"`
	Products []Product	
}