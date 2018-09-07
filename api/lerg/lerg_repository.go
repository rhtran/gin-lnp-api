package lerg

import (
	"github.com/gocql/gocql"
	"fmt"
	)

// LergRepository retrieves lerg from the database
type LergRepository struct {
	session *gocql.Session
}

// NewLergRepository creates a new LergRepository
func NewLergRepository(session *gocql.Session ) *LergRepository {
	return &LergRepository{session: session}
}

// Get reads the lerg with the specified ID from the database.
func (repository *LergRepository) FindByNpaNxx(npa string, nxx string, blockId string) (*Lerg, error) {
	var lerg Lerg

	// Compose the CQL query.
	query := fmt.Sprintf("SELECT country_code, npa, nxx, block_id, tbp_ind, lata, ltype, rc, state, " +
		"zip, zip2, zip3, zip4, country, tz, dst, longitude, latitude, olson, utc, ocn, type, company, " +
		"common_name, last_updated " +
		"FROM npa_nxx WHERE country_code = '1' and npa = '%s' " +
		"and nxx = '%s' and block_id = '%s' LIMIT 1", npa, nxx, blockId)

	q := repository.session.Query(query)

	var zips [4]string

	err := q.Scan(&lerg.CountryCode, &lerg.Npa, &lerg.Nxx, &lerg.BlockId, &lerg.TbpInd, &lerg.Lata, &lerg.Ltype,
		&lerg.Rc, &lerg.State, &zips[0], &zips[1], &zips[2], &zips[3], &lerg.Country, &lerg.Tz,
		&lerg.Dst, &lerg.Longitude, &lerg.Latitude, &lerg.Olson, &lerg.Utc, &lerg.Ocn, &lerg.Type,
		&lerg.Company, &lerg.CommonName, &lerg.LastUpdated)

	if err == nil {
		for _, str := range zips {
			if str != "" {
				lerg.Zips = append(lerg.Zips, str)
			}
		}
	}

	return &lerg, err
}