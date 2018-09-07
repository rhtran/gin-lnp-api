package ocn

// Ocn represents an ocn record.
type Ocn struct {
	Ocn string			`json:"ocn" db:"ocn"`
	Type string			`json:"type" db:"type"`
	Neca string			`json:"neca" db:"neca"`
	Company string		`json:"company" db:"company"`
	CommonName string	`json:"common_name" db:"common_name"`
}