package internal

import (
	"fmt"
	"testing"
)

func TestStoragePut(t *testing.T) {
	s := NewStorage("test.yaml")

	if err := s.Put([]byte("test")); err != nil {
		t.Errorf("put failed: %v", err)
	}
}

func TestStorageGet(t *testing.T) {
	s := NewStorage("test.yaml")

	data, err := s.Get()
	if err != nil {
		t.Errorf("get failed: %v", err)
	}

	fmt.Println(string(data))
}
