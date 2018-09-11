package ocn

import (
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFindByOcn(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error ‘%s’ was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string {
		"ocn", "neca", "type", "common_name", "company"}).
		AddRow("874E", "N", "L_RESELLER", "Points South", "Points South")
	query := "SELECT ocn, neca, type, common_name, company FROM ocn WHERE ocn = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := NewOcnRepository(db)
}