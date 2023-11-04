package models

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Title       string
	Description string
	Photo       string
	Steps       []RecipeStep
	Ingredients []Ingredient
	AuthorID    uint
	Favorites   []User `gorm:"many2many:favorites;"`
}

type RecipeStep struct {
	gorm.Model
	RecipeID uint
	Step     string
	//include descriptions and photo fields here
}

type Ingredient struct {
	gorm.Model
	RecipeID uint
	Name     string
	Amount   string
}
