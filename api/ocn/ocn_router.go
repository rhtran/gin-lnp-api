package ocn

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OcnRouter struct {
	ocnService *OcnService
}

func NewOcnRouter(ocnService *OcnService) *OcnRouter {
	return &OcnRouter{
		ocnService: ocnService,
	}
}

func (ocnRouter *OcnRouter) OcnRegister(router *gin.RouterGroup) {
	router.GET(":id", ocnRouter.OcnRetrieve)
}

func (ocnRouter *OcnRouter) OcnRetrieve(c *gin.Context) {
	id := c.Param("id")
	ocn, err := ocnRouter.ocnService.GetByOcn(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, ocn)
}
