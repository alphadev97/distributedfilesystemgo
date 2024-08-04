package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "mamasbestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "4018a9a4e2c66eae1c409b5e46a761001d50f15f"
	expectedPathName := "4018a/9a4e2/c66ea/e1c40/9b5e4/6a761/001d5/0f15f"

	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Original != expectedPathName {
		t.Errorf("have %s want %s", pathKey.Original, expectedOriginalKey)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}

}
