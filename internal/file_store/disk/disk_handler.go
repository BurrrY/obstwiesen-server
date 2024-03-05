package disk

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/BurrrY/obstwiesen-server/graph/model"
	"github.com/BurrrY/obstwiesen-server/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
)

type stor struct {
	ConnectionError error
	BasePath        string
	SetupDone       bool
}

func (s stor) Setup() {

	if s.SetupDone == true {
		return
	}

	if viper.GetString(config.FILE_PROVIDER) != "disk" {
		log.New().Info("Skip disk Init by Config: " + viper.GetString(config.FILE_PROVIDER))
		Thing.ConnectionError = errors.New("disk disabled")
		return
	}

	path := viper.GetString(config.FILE_CONNSTR)
	Thing.BasePath = path

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

	s.SetupDone = true
}

var Thing stor

func (s stor) GetType() string {
	return "disk"
}

func (s stor) GetFiles(parentId string) ([]*model.File, error) {

	base_path := viper.GetString(config.PUBLIC_URL)

	res := []*model.File{}
	entries, err := os.ReadDir(filepath.Join(Thing.BasePath, parentId))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			//fmt.Println("The file does not exist.")
			return res, nil
		} else {
			log.Error("GetFiles ", err.Error())
			return res, nil
		}
	}

	for _, e := range entries {
		res = append(res, &model.File{
			ParentID: parentId,
			Path:     base_path + "/assets/" + parentId + "/" + e.Name(),
		})
	}

	return res, nil
}

func (s stor) StoreFile(file *graphql.Upload, parentID string, fileID string) error {

	log.Info("Store File:", file.Filename)
	log.Info("Target:", filepath.Join(Thing.BasePath, parentID, fileID))

	err := os.MkdirAll(filepath.Join(Thing.BasePath, parentID), os.ModePerm)
	if err != nil {
		log.Error(err)
		return err
	}

	f2, err := os.Create(filepath.Join(Thing.BasePath, parentID, fileID))
	if err != nil {
		log.Error(err)
	}
	defer f2.Close()
	io.Copy(f2, file.File)

	return nil
}
