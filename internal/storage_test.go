package internal

import (
	"fmt"
	"testing"
)

func TestStoragePut(t *testing.T) {
	s := newStorage("test.yaml")

	if err := s.put([]byte("test")); err != nil {
		t.Errorf("put failed: %v", err)
	}
}

func TestStorageGet(t *testing.T) {
	s := newStorage("test.yaml")

	data, err := s.get()
	if err != nil {
		t.Errorf("get failed: %v", err)
	}

	fmt.Println(string(data))
}
