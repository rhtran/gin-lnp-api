package ocn


func newMockOcnRepository() ocnRepository {
	return &mockOcnRepository{
		records: []Ocn{
			{Ocn: "3912", Neca: "N", Type: "CLEC", CommonName: "Telscape Communications", Company: "Telscape Communications Inc"},
			{Ocn: "539G", Neca: "N", Type: "CLEC", CommonName: "Aeon Communications", Company: "Aeon Communications LLC - TX"},
			{Ocn: "955F", Neca: "N", Type: "WIRELESS", CommonName: "Telecom North America Mobile", Company: "Telecom North America Mobile Inc"},
			{Ocn: "7309", Neca: "N", Type: "CLEC", CommonName: "AT&T", Company: "TCG New Jersey Inc - PA"},
		},
	}
}

type mockOcnRepository struct {
	records []Ocn
}