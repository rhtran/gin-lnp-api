package main

import (
	"gin-lnp-api/api/ocn"
	"gin-lnp-api/api/lerg"
	"gin-lnp-api/api/lrn"
	"github.com/gin-gonic/gin"
		"fmt"
	"gin-lnp-api/dbase"
	"gin-lnp-api/app"
	"gin-lnp-api/ds"
)

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

	httpServer := ds.NewHttpServer(r)
	httpServer.Start()
}
