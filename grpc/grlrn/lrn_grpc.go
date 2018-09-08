package grlrn

import (
	"golang.org/x/net/context"
	"gin-lnp-api/api/lrn"
)

type GrpcLrnService struct {
	lrnService *lrn.LrnService
}

func NewGrpcLrnService(lrnService *lrn.LrnService) *GrpcLrnService {
	return &GrpcLrnService {
		lrnService: lrnService,
	}
}

func (s *GrpcLrnService) RetrieveByDid(ctx context.Context, didReq *DidReq) (*LrnRes, error) {
	lrn, err := s.lrnService.GetByDid(didReq.Did)
	if err != nil {
		return nil, err
	}

	return &LrnRes{
		Did: lrn.Did,
		Lrn: lrn.Lrn,
		Ocn: *lrn.Ocn,
		Ptype: *lrn.Type,
		Company: *lrn.Company,
		Bname: *lrn.CommonName,
		Dst: *lrn.Dst,
		City: *lrn.Rc,
		St: *lrn.State,
		//Zips: *lrn.Zips,
		Ct: *lrn.Country,
		Tz: *lrn.Tz,
		Lat: *lrn.Latitude,
		Long: *lrn.Longitude,
		Lnp: int64(lrn.Lnp),
		//Lu: lrn.LastUpdated,
	}, nil
}

func (s *GrpcLrnService) FindByDids(ctx context.Context, lrn *DidsReq) (*LrnsRes, error) {
	return &LrnsRes{}, nil
}

func (s *GrpcLrnService) FindByStreamDid(stream LrnService_FindByStreamDidServer) error {
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
