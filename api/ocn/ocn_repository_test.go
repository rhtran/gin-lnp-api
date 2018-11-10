package ocn

import (
	"testing"
	"github.com/golang/mock/gomock"
)

func  TestFindByOcn(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOcnRepository := NewMockocnRepository(mockCtrl)
	testOcn := Ocn{Ocn: "9962", CommonName: "Winstar Communications", Company: "Winstar Communications LLC - CA",
					Type: "CLEC", Neca:"N"}
	mockOcnRepository.EXPECT().FindByOcn("9962").Return(&testOcn, nil).Times(1)
	testGetByOcn(t, mockOcnRepository)
}

func testGetByOcn(t *testing.T, mockrepository ocnRepository) {
	testOcn := Ocn{Ocn: "9962", CommonName: "Winstar Communications", Company: "Winstar Communications LLC - CA",
		Type: "CLEC", Neca:"N"}

	ocnService := NewOcnService(mockrepository)

	r, err := ocnService.GetByOcn("9962")

	if err != nil || *r != testOcn {
		t.Errorf("findByOcn failed")
	}
	t.Log("Reply : ", r)
}