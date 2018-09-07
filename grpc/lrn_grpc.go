package grpc

import (
	"golang.org/x/net/context"
)

type GrpcLrnServer struct {

}

func (s *GrpcLrnServer) FindByDid(ctx context.Context, lrn *DidReq) (*LrnRes, error) {
	return &LrnRes{}, nil
}

func (s *GrpcLrnServer) FindByDids(ctx context.Context, lrn *DidsReq) (*LrnsRes, error) {
	return &LrnsRes{}, nil
}

func (s *GrpcLrnServer) FindByStreamDid(stream LrnService_FindByStreamDidServer) error {
	for {
		_, err := stream.Recv()
		//if err == io.EOF {
		//	return nil
		//}
		//if err != nil {
		//	return err
		//}
		//key := serialize(in.DidReq)
		//// look for notes to be sent to client
		//
		//for _, note := range s.[key] {
		//	if err := stream.Send(note); err != nil {
		//		return err
		//	}
		//}
		return err
	}
}
