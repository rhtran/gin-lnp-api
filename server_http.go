package main

import (
	"github.com/gin-gonic/gin"
	"gin-lnp-api/api/ocn"
	"gin-lnp-api/dbase"
	"fmt"
	"gin-lnp-api/app"
	"gin-lnp-api/api/lerg"
	"gin-lnp-api/api/lrn"
)

//type HttpServer struct {
//	lergService *lerg.LergService
//	lrnService *lrn.LrnService
//	ocnService *ocn.OcnService
//}
//
//func NewHttpServer(lergService *lerg.LergService, lrnService *lrn.LrnService, ocnService *ocn.OcnService) *HttpServer {
//	return &HttpServer{
//		lergService: lergService,
//		lrnService: lrnService,
//		ocnService: ocnService,
//	}
//}


func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	db := dbase.ConnectDatabase()
	defer db.Close()

	// OCN
	ocnRepository := ocn.NewOcnRepository(db)
	ocnService := ocn.NewOcnService(ocnRepository)
	ocnRouter := ocn.NewOcnRouter(ocnService)

	//LERG
	lergRepository := lerg.NewLergRepository(db)
	lergService := lerg.NewLergService(lergRepository)
	lergRouter := lerg.NewLergRouter(lergService)

	// LRN
	lrnRepository := lrn.NewLrnRepository(db)
	lrnService := lrn.NewLrnService(lergRepository, lrnRepository)
	lrnRouter := lrn.NewLrnRouter(lrnService)


	r := gin.Default()
	v1 := r.Group("/v1/lnp")

	ocnRouter.OcnRegister(v1.Group("/ocns"))
	lergRouter.LergRegister(v1.Group("/lergs"))
	lrnRouter.LrnRegister(v1.Group("/dids"))

	r.Run()
}