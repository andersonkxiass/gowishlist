//+build wireinject

package main

import (
	. "WishlistApi/repositories"
	. "WishlistApi/services"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initWishAPI(db *gorm.DB) WishAPI {
	wire.Build(ProvideWishRepository, ProvideWishAPI)

	return WishAPI{}
}

func initListAPI(db *gorm.DB) ListAPI {
	wire.Build(ProvideListRepository, ProvideListAPI)

	return ListAPI{}
}
