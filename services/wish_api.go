package services

import (
	. "WishlistApi/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

import (
	_ "github.com/gin-gonic/gin"
	_ "log"
	_ "net/http"
	_ "strconv"
)

type WishAPI struct {
	WishRepository WishRepository
}

func ProvideWishAPI(p WishRepository) WishAPI {
	return WishAPI{WishRepository: p}
}

func (p *WishAPI) FindAll(c *gin.Context) {
	wishDB := p.WishRepository.FindAll()
	c.JSON(http.StatusOK, gin.H{"data": wishDB})
}
