package storage

import (
	"github.com/BurrrY/obstwiesen-server/internal/data_store/sqlite"
	log "github.com/sirupsen/logrus"
	"os"
)

func GetProvider() (*Storage, error) {
	var result Storage

	dbProvider := os.Getenv("PROVIDER")

	if dbProvider == "sqlite" {
		result = &sqlite.Connection
	} else {
		log.WithFields(log.Fields{
			"connectionData": dbProvider,
		}).Fatal("Cannot Connect to DB, No 'dbProvider' configured!")
	}

	return &result, nil
}
