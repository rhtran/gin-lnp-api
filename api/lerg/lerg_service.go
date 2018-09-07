package lerg


// LergService provides services related with lergs.
type LergService struct {
	repository *LergRepository
}

// NewLergService creates a new LergService with the given lerg DAO.
func NewLergService(repository *LergRepository) *LergService {
	return &LergService{repository: repository}
}

// Get returns the lerg with the specified the lerg ID.
func (service *LergService) GetByNpaNxx(npa string, nxx string, blockId string) (*Lerg, error) {
	var lerg *Lerg
	var err error
	lerg, err = service.repository.FindByNpaNxx(npa, nxx, blockId)

	if err != nil {
		lerg, err = service.repository.FindByNpaNxx(npa, nxx, "A")
	}

	return lerg, err
}