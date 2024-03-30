package schemas

import "gorm.io/gorm"

type Gorga struct {
	gorm.Model
	Id       int64
	Name     string
	Role     string
	Employed bool
}
