package ocn

type ocnService interface {
	GetByOcn(ocn string) (*Ocn, error)
}

type OcnService struct {
	repository ocnRepository
}

// NewArtistService creates a new ArtistService with the given artist DAO.
func NewOcnService(repository ocnRepository) *OcnService {
	return &OcnService{repository: repository}
}

// Get returns the artist with the specified the artist ID.
func (service *OcnService) GetByOcn(ocn string) (*Ocn, error) {
	return service.repository.FindByOcn(ocn)
}