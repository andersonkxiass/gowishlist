package repositories

import (
	. "WishlistApi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


type WishRepository struct {
	DB *gorm.DB
}

func ProvideWishRepository(DB *gorm.DB) WishRepository {
	return WishRepository{DB: DB}
}

func (p *WishRepository) FindAll() []Wish {
	var products []Wish
	p.DB.Debug().Find(&products)

	return products
}