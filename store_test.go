package main

import (
	"bytes"
	"testing"
)

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}

	s := NewStore(opts)
	data := bytes.NewReader([]byte("some data "))
	if err := s.writeStream("mytake1", data); err != nil {
		t.Error(err)
	}
}
