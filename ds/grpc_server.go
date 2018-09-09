package ds

import (
	"google.golang.org/grpc/reflection"
	"net"
	"log"
	hpb "gin-lnp-api/grpc/grhello"
	"google.golang.org/grpc"
				)

type GrpcServer struct {
	s *grpc.Server
	lis net.Listener
}

func NewGrpcServer(s *grpc.Server, lis net.Listener) *GrpcServer {
	return &GrpcServer{
		s: s,
		lis: lis,
	}
}


func (grpc *GrpcServer) Start() {
	helloSrv := hpb.NewGrpcHelloService()
	//ocnSrv := ocnpb.NewGrpcOcnService(ocnService)
	//lrnSrv := lrnpb.NewGrpcLrnService(lrnService)

	hpb.RegisterHelloServiceServer(grpc.s, helloSrv)
	//ocnpb.RegisterOcnServiceServer(s, ocnSrv)
	//lrnpb.RegisterLrnServiceServer(s, lrnSrv)
	log.Println("Start Grpc Server ...")

	// Register reflection service on gRPC server.
	reflection.Register(grpc.s)
	if err := grpc.s.Serve(grpc.lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//log.Println("Shutdown Http Server ...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := grpc.lis.Close(); err != nil {
	//	log.Fatal("Grpc Server Shutdown:", err)
	//}

	log.Println("Grpc Server exiting")
}