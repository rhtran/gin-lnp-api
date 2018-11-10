package ds

import (
	"google.golang.org/grpc/reflection"
		"log"
		"google.golang.org/grpc"
				)

type GrpcServer struct {
	s *grpc.Server
}

func NewGrpcServer(s *grpc.Server) *GrpcServer {
	return &GrpcServer{
		s: s,
	}
}


func (grpc *GrpcServer) GrpcRouter() *grpc.Server {
	//helloSrv := hpb.NewGrpcHelloService()

	//lrnSrv := lrnpb.NewGrpcLrnService(lrnService)

	//hpb.RegisterHelloServiceServer(grpc.s, helloSrv)
	//lrnpb.RegisterLrnServiceServer(s, lrnSrv)
	log.Println("Start Grpc Server ...")

	// Register reflection service on gRPC server.
	reflection.Register(grpc.s)

	return grpc.s
}