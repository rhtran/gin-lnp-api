package lrn

import (
	"github.com/gocql/gocql"
	"fmt"
)

// LrnRepository retrieves lrn from the database
type LrnRepository struct {
	session *gocql.Session
}

// NewLrnRepository creates a new LrnRepository
func NewLrnRepository(session *gocql.Session ) *LrnRepository {
	return &LrnRepository{session: session}
}

// Get reads the lrn with the specified ID from the database.
func (repository *LrnRepository) FindByDid(did string) (*Lrn, error) {
	var lrn Lrn

	// Compose the CQL query.
	query := fmt.Sprintf("SELECT did, lrn, ocn, type, company, common_name, rc, state, country, tz, dst, " +
		"zip, zip2, zip3, zip4, longitude, latitude, last_updated " +
		"FROM lrn WHERE country_code = '1' and did = '%s'", did)

	var zips [4]string

	q := repository.session.Query(query)

	err := q.Scan(&lrn.Did, &lrn.Lrn, &lrn.Ocn, &lrn.Type, &lrn.Company, &lrn.CommonName, &lrn.Rc, &lrn.State, &lrn.Country,
		&lrn.Tz, &lrn.Dst, &zips[0], &zips[1], &zips[2], &zips[3], &lrn.Longitude, &lrn.Latitude, &lrn.LastUpdated)

	if err == nil {
		lrn.Lnp = 1
		lrn.Zips = make([]string, 0)

		for _, str := range zips {
			if str != "" {
				lrn.Zips = append(lrn.Zips, str)
			}
		}
	}

	return &lrn, err
}