package internal

import (
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

type storage struct {
	path string
}

func newStorage(filename string) *storage {
	dir, _ := homedir.Dir()

	return &storage{
		path: path.Join(dir, storageDir, filename),
	}
}

func (s *storage) get() ([]byte, error) {
	return os.ReadFile(s.path)
}

func (s *storage) put(data []byte) error {
	dir := path.Dir(s.path)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err := os.Mkdir(dir, 0777)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(s.path, data, 0777)
}
