package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
	hpb "gin-lnp-api/grpc/grhello"
				"gin-lnp-api/app"
	"fmt"
	"gin-lnp-api/dbase"
		)

func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	db := dbase.ConnectDatabase()
	defer db.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	helloSrv := hpb.NewGrpcHelloService()

	// OCN
	// OCN
	//ocnRepository := ocn.NewOcnRepository(db)
	//ocnService := ocn.NewOcnService(ocnRepository)
	//ocnSrv := ocnpb.NewGrpcOcnService(ocnService)

	// LRN
	//lergRepository := lerg.NewLergRepository(db)
	//lrnRepository := lrn.NewLrnRepository(db)
	//lrnService := lrn.NewLrnService(lergRepository, lrnRepository)
	//lrnSrv := lrnpb.NewGrpcLrnService(lrnService)

	hpb.RegisterHelloServiceServer(s, helloSrv)
	//ocnpb.RegisterOcnServiceServer(s, ocnSrv)
	//lrnpb.RegisterLrnServiceServer(s, lrnSrv)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}