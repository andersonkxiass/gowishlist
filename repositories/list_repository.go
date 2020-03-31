package repositories

import (
	. "WishlistApi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


type ListRepository struct {
	DB *gorm.DB
}

func ProvideListRepository(DB *gorm.DB) ListRepository {
	return ListRepository{DB: DB}
}

func (p *ListRepository) FindAll() []List {
	var list []List
	p.DB.Debug().Preload("Wish").Find(&list)

	return list
}
