package kv

import (
	"os"

	"github.com/lotusdblabs/lotusdb/v2"
)

type lotusdbStore struct {
	db           *lotusdb.DB
	writeOptions *lotusdb.WriteOptions
}

func newLotusdb(path string) (Store, error) {
	options := lotusdb.DefaultOptions
	if len(path) != 0 {
		_ = os.MkdirAll(path, os.ModePerm)
	}
	db, err := lotusdb.Open(options)
	if err != nil {
		return nil, err
	}
	return &lotusdbStore{db: db, writeOptions: &lotusdb.WriteOptions{Sync: false, DisableWal: false}}, nil
}

func (s *lotusdbStore) Put(key []byte, value []byte) error {
	return s.db.Put(key, value, s.writeOptions)
}

func (s *lotusdbStore) Get(key []byte) ([]byte, error) {
	return s.db.Get(key)
}

func (s *lotusdbStore) Delete(key []byte) error {
	return s.db.Delete(key, s.writeOptions)
}

func (s *lotusdbStore) Close() error {
	return s.db.Close()
}
