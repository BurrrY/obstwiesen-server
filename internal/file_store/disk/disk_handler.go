package disk

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/BurrrY/obstwiesen-server/graph/model"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

type stor struct {
	ConnectionError error
	BasePath        string
}

var Handler stor

func init() {

	if os.Getenv("FILE_PROVIDER") != "disk" {
		log.New().Info("Skip disk Init by Config: " + os.Getenv("FILE_PROVIDER"))
		Handler.ConnectionError = errors.New("disk disabled")
		return
	}

	path := os.Getenv("FILE_CONNSTR")
	Handler.BasePath = path

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

}

func (s stor) GetType() string {
	return "disk"
}

func (s stor) GetFiles(parentId string) ([]*model.File, error) {

	res := []*model.File{}
	entries, err := os.ReadDir(filepath.Join(Handler.BasePath, parentId))
	if err != nil {
		log.Error(err)
		return res, nil
	}

	for _, e := range entries {
		res = append(res, &model.File{
			ParentID: parentId,
			Path:     "/assets/" + parentId + "/" + e.Name(),
		})
	}

	return res, nil
}

func (s stor) StoreFile(file *graphql.Upload, parentID string, fileID string) error {

	log.Info("Store File:", file.Filename)
	log.Info("Target:", filepath.Join(Handler.BasePath, parentID, fileID))

	err := os.MkdirAll(filepath.Join(Handler.BasePath, parentID), os.ModePerm)
	if err != nil {
		log.Error(err)
		return err
	}

	f2, err := os.Create(filepath.Join(Handler.BasePath, parentID, fileID))
	if err != nil {
		log.Error(err)
	}
	defer f2.Close()
	io.Copy(f2, file.File)

	return nil
}
