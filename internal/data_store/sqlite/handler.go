package sqlite

import (
	"errors"
	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
)

type stor struct {
	ConnectionString string
	ConnectionError  error
}

var db *sqlx.DB
var Connection stor

func (m stor) GetType() string {
	return "sqlite"
}

func init() {

	if os.Getenv("PROVIDER") != "sqlite" {
		log.New().Info("Skip sqlite Init by Config: " + os.Getenv("PROVIDER"))
		Connection.ConnectionError = errors.New("sqlite disabled")
		return
	}

	var err error

	db, err = sqlx.Connect("sqlite3", os.Getenv("CON_STR"))
	if err != nil {
		log.Fatalln(err)
	}

	updateDb()

}
