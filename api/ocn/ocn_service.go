package ocn

type OcnService struct {
	repository *OcnRepository
}

// NewArtistService creates a new ArtistService with the given artist DAO.
func NewOcnService(repository *OcnRepository) *OcnService {
	return &OcnService{repository: repository}
}

// Get returns the artist with the specified the artist ID.
func (service *OcnService) GetByOcn(ocn string) (*Ocn, error) {
	return service.repository.FindByOcn(ocn)
}