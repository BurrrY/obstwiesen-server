package storage

import (
	"github.com/BurrrY/obstwiesen-server/internal/file_store/disk"
	log "github.com/sirupsen/logrus"
	"os"
)

func GetProvider() (*FileStorage, error) {
	var result FileStorage

	dbProvider := os.Getenv("FILE_PROVIDER")

	if dbProvider == "disk" {
		result = &disk.Handler
	} else {
		log.WithFields(log.Fields{
			"connectionData": dbProvider,
		}).Fatal("Cannot Connect to FILE_PROVIDER, No 'FILE_PROVIDER' configured!")
	}

	return &result, nil
}
