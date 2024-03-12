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

// Get reads the data from the file
func (s *Storage) Get() ([]byte, error) {
	err := ensureFileExists(s.path)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(s.path)
}

// Put writes the data to the file
func (s *Storage) Put(data []byte) error {
	err := ensureFileExists(s.path)
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0755)
}

// ensureFileExists ensures that the directory and file exists
func ensureFileExists(filepath string) error {
	dir := path.Dir(filepath)

	// ensure the directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// ensure the file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		_, err := os.Create(filepath)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
