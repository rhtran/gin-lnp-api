package ocn

import (
	"github.com/gocql/gocql"
	"fmt"
)

type ocnRepository interface {
	FindByOcn(ocn string) (*Ocn, error)
}

// OcnRepository retrieves ocn from the database
type OcnRepository struct {
	session *gocql.Session
}

// NewOcnRepository creates a new OcnRepository
func NewOcnRepository(session *gocql.Session ) *OcnRepository {
	return &OcnRepository{session: session}
}

// Get reads the ocn with the specified ID from the database.
func (repository *OcnRepository) FindByOcn(id string) (*Ocn, error) {
	var ocn Ocn

	// Compose the CQL query.
	query := fmt.Sprintf("SELECT ocn, type, neca, company, common_name FROM ocn WHERE ocn = '%s' LIMIT 1", id)

	q := repository.session.Query(query)
	err := q.Scan(&ocn.Ocn, &ocn.Type, &ocn.Neca, &ocn.Company, &ocn.CommonName)

	return &ocn, err
}


