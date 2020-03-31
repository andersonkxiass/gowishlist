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

type ListAPI struct {
	ListRepository ListRepository
}

func ProvideListAPI(p ListRepository) ListAPI {
	return ListAPI{ListRepository: p}
}

func (p *ListAPI) ListALl(c *gin.Context) {
	wishDB := p.ListRepository.FindAll()
	c.JSON(http.StatusOK, gin.H{"data": wishDB})
}
