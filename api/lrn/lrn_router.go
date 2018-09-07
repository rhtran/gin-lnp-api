package lrn

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type LrnRouter struct {
	lrnService *LrnService
}

func NewLrnRouter(lrnService *LrnService) *LrnRouter {
	return &LrnRouter{
		lrnService: lrnService,
	}
}

func (lrnRouter *LrnRouter) LrnRegister(router *gin.RouterGroup) {
	router.GET(":did", lrnRouter.LrnRetrieve)
	router.GET("", lrnRouter.LrnsRetrieve)
}


func (lrnRouter *LrnRouter) LrnRetrieve(c *gin.Context) {
	did := c.Param("did")
	lrn, err := lrnRouter.lrnService.GetByDid(did)

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"": lrn})
}

func (lrnRouter *LrnRouter) LrnsRetrieve(c *gin.Context) {
	didQuery := c.Query("did_list")

	rp := regexp.MustCompile(",")
	didList := rp.Split(didQuery, -1)
	responses, err := lrnRouter.lrnService.GetByDids(didList)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"": responses})
}