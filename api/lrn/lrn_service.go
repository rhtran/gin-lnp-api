package lrn

import (
	"gin-lnp-api/api/lerg"
	"regexp"
	"strings"
	"time"
	"sync"
)

type LrnService struct {
	lergRepository *lerg.LergRepository
	lrnRepository *LrnRepository
}

// NewArtistService creates a new ArtistService with the given artist DAO.
func NewLrnService(lergRepository *lerg.LergRepository, lrnRepository *LrnRepository) *LrnService {
	return &LrnService{lergRepository: lergRepository,
						lrnRepository: lrnRepository}
}

func RemoveCountryCode(did string) string {
	re := regexp.MustCompile("[^\\d]")
	did = strings.Replace(did, re.FindString(did), "", 1)

	if strings.HasPrefix(did, "1") {
		did = strings.Replace(did, "1", "", 1)
	}

	return did
}

func IsDidValid(did string) bool {
	re := regexp.MustCompile("^\\d{10}$")
	return re.MatchString(did)
}

func (s *LrnService) getLrnInfo(did string) (*Lrn, error)  {
	npa := did[0:3]
	nxx := did[3:6]
	blockId := did[6:7]

	var lrg *lerg.Lerg
	var err error

	lrg, err = s.lergRepository.FindByNpaNxx(npa, nxx, blockId)

	if err != nil {
		lrg, err = s.lergRepository.FindByNpaNxx(npa, nxx, "A")

		if err != nil {
			*lrg = lerg.Lerg{}
			lrg.Zips = make([]string, 0)
		}
	}

	return toLrn(lrg, did), nil

}

/**
 * Map Lerg to Lrn
 */
func toLrn(lerg *lerg.Lerg, did string) (*Lrn) {
	return &Lrn {
		Did: did,
		Lrn: did,
		Ocn: &lerg.Ocn,
		Type: &lerg.Type,
		Company: &lerg.Company,
		CommonName: &lerg.CommonName,
		Rc: &lerg.Rc,
		State: &lerg.State,
		Zips: lerg.Zips,
		Country: &lerg.Country,
		Tz: &lerg.Tz,
		Dst: &lerg.Dst,
		Latitude: &lerg.Latitude,
		Longitude: &lerg.Longitude,
		Lnp: 0,
	}
}

// Get returns the lrn with the specified the lrn ID.
func (s *LrnService) GetByDid(did string) (*Lrn, error) {
	//rs.Infof("Retrieve did: %v", did)

	did = RemoveCountryCode(did)
	var lrn *Lrn
	var err error

	if IsDidValid(did) {
		lrn, err = s.lrnRepository.FindByDid(did)

		if err != nil {
			lrn, err = s.getLrnInfo(did)
		}

		return lrn, err
	} else {
		lrn = &Lrn{}
		lrn.Did = did
		lrn.Lrn = did
		lrn.Zips = make([]string, 0)
		lrn.LastUpdated = time.Now()

		return lrn, nil
	}
}

//Concurrency get list of dids
func ChanGet(s *LrnService, did string, lrns *[]Lrn, wg *sync.WaitGroup) {
	defer wg.Done()

	lrn, _ := s.GetByDid(did)
	*lrns = append(*lrns, *lrn)
}

// Query retrieves the lrn records with the specified list of dids.
func (s *LrnService) GetByDids(didList []string) ([]Lrn, error) {
	arrLen := len(didList)
	var wg sync.WaitGroup

	//rs.Info(fmt.Sprintf("Retrieve did list: %v", didList))
	// create array with size and cap = length of did list.
	lrns := make([]Lrn, 0, arrLen)

	wg.Add(arrLen)
	for i := 0; i < arrLen; i++ {
		go ChanGet(s, didList[i], &lrns, &wg)
	}
	wg.Wait()

	return lrns, nil
}
