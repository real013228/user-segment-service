package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Segments []Segment `gorm:"many2many:user_segments;"`
}
