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

func NewGrpcServer(s *grpc.Server) *GrpcServer {
	return &GrpcServer{
		s: s,
	}
}


func (grpc *GrpcServer) GrpcRouter() *grpc.Server {
	helloSrv := hpb.NewGrpcHelloService()
	//ocnSrv := ocnpb.NewGrpcOcnService(ocnService)
	//lrnSrv := lrnpb.NewGrpcLrnService(lrnService)

	hpb.RegisterHelloServiceServer(grpc.s, helloSrv)
	//ocnpb.RegisterOcnServiceServer(s, ocnSrv)
	//lrnpb.RegisterLrnServiceServer(s, lrnSrv)
	log.Println("Start Grpc Server ...")

	// Register reflection service on gRPC server.
	reflection.Register(grpc.s)

	return grpc.s
}