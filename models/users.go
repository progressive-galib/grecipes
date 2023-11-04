package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Email     string
	Password  string   // Consider using a hashed password
	Recipes   []Recipe `gorm:"foreignKey:AuthorID"`
	Favorites []Recipe `gorm:"many2many:favorites;"`
}

type Favorites struct {
	gorm.Model
	UserID   uint
	RecipeID uint
}
