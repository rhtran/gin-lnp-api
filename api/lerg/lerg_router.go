package lerg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LergRouter struct {
	lergService *LergService
}

func NewLergRouter(lergService *LergService) *LergRouter {
	return &LergRouter {
		lergService: lergService,
	}
}

func (lergRouter *LergRouter) LergRegister(router *gin.RouterGroup) {
	router.GET(":npa/:nxx/:block_id", lergRouter.LergRetrieve)
}

func (lergRouter *LergRouter) LergRetrieve(c *gin.Context) {
	npa := c.Param("npa")
	nxx := c.Param("nxx")
	blockId := c.Param("block_id")

	response, err := lergRouter.lergService.GetByNpaNxx(npa, nxx, blockId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"": response})
}