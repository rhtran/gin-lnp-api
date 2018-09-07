package main

import (
	"gin-lnp-api/dbase"
	"fmt"
	"gin-lnp-api/app"
	"gin-lnp-api/api/ocn"
	"log"
)

func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// load error messages
	//if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
	//	panic(fmt.Errorf("failed to read the error message file: %s", err))
	//}

	// create the logger
	//logger := logrus.New()
	db := dbase.ConnectDatabase()
	defer db.Close()

	ocnRepository := ocn.NewOcnRepository(db)
	ocnService := ocn.NewOcnService(ocnRepository)

	ocn, err := ocnService.GetByOcn("348H")

	if err != nil {
		log.Fatalf("Error retrieving ocn: %v", err)
	}

	fmt.Printf("%s: %s, %s, %s", ocn.Ocn, ocn.Company, ocn.CommonName, ocn.Type)




	//r.Run() // listen and server on 0.0.0.0:8080
}
