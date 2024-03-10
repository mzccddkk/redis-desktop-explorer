package internal

import (
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

type Storage struct {
	path string
}

func NewStorage(filename string) *Storage {
	dir, _ := homedir.Dir()

	return &Storage{
		path: path.Join(dir, storageDir, filename),
	}
}

func (s *Storage) Get() ([]byte, error) {
	return os.ReadFile(s.path)
}

func (s *Storage) Put(data []byte) error {
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
