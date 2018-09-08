package grocn

import (
	"gin-lnp-api/api/ocn"
	"context"
)

type GrpcOcnService struct {
	ocnService *ocn.OcnService
}

func NewGrpcOcnService(ocnService *ocn.OcnService) *GrpcOcnService {
	return &GrpcOcnService {
		ocnService: ocnService,
	}
}

func (s *GrpcOcnService) Greet(ctx context.Context, ocnReq *OcnReq) (*OcnRes, error) {
	ocn, err := s.ocnService.GetByOcn(ocnReq.GetOcn())
	if err != nil {
		return nil, err
	}

	return &OcnRes{
		Ocn: ocn.Ocn,
		Type: ocn.Type,
		Neca: ocn.Neca,
		Company: ocn.Company,
		CommonName: ocn.CommonName,
	}, nil
}
