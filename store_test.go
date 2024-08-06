package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "mamasbestpicture"
	pathKey := CASPathTransformFunc(key)
	exptectedFileName := "4018a9a4e2c66eae1c409b5e46a761001d50f15f"
	expectedPathName := "4018a/9a4e2/c66ea/e1c40/9b5e4/6a761/001d5/0f15f"

	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Filename != exptectedFileName {
		t.Errorf("have %s want %s", pathKey.Filename, exptectedFileName)
	}
}

func TestStore(t *testing.T) {

	s := newStore()
	defer teardown(t, s)

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("some jpg bytes")

		if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); !ok {
			t.Errorf("exptected to have key %s", key)
		}

		r, err := s.Read(key)
		if err != nil {
			t.Error(err)
		}

		b, _ := ioutil.ReadAll(r)
		if string(b) != string(data) {
			t.Errorf("want %s have %s", data, b)
		}

		fmt.Println(string(b))

		if err := s.Delete(key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); ok {
			t.Errorf("exptected to NOT have key %s", key)
		}
	}

}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
