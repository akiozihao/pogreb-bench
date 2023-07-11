package kv

import (
	"os"

	"github.com/rosedblabs/rosedb/v2"
)

type rosedbStore struct {
	db *rosedb.DB
}

func newRosedb(path string) (Store, error) {
	options := rosedb.DefaultOptions
	if len(path) != 0 {
		_ = os.MkdirAll(path, os.ModePerm)
		options.DirPath = path
	}
	db, err := rosedb.Open(rosedb.DefaultOptions)
	if err != nil {
		return nil, err
	}
	return &rosedbStore{db: db}, nil
}

func (s *rosedbStore) Put(key []byte, value []byte) error {
	return s.db.Put(key, value)
}

func (s *rosedbStore) Get(key []byte) ([]byte, error) {
	var val []byte
	val, err := s.db.Get(key)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (s *rosedbStore) Delete(key []byte) error {
	return s.db.Delete(key)
}

func (s *rosedbStore) Close() error {
	return s.db.Close()
}
