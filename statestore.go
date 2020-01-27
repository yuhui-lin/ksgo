package ksgo

import (
	"github.com/dgraph-io/badger"
)

type StateStore interface {
	// Name() string
	// Open() error
	Close() error
}

type ReadOnlyKVStore interface {
	StateStore
	Has(string) (bool, error)
	Get(string) ([]byte, error)
	GetOffset(defValue int64) (int64, error)
	Iterator() (Iterator, error)
	IteratorWithRange(start, limit []byte) (Iterator, error)
	// MarkRecovered() error
	// Recovered() bool
}

type KVStore interface {
	ReadOnlyKVStore
	Set(string, []byte) error
	Delete(string) error
	// SetOffset(value int64) error
}

type badgerKVStore struct {
	db *badger.DB
}

func NewBadgerKVStore(opt badger.Options) *badgerKVStore {
	db, err := badger.Open(opt)
	if err != nil {
		panic(err)
	}
	return &badgerKVStore{db: db}
}

func (s *badgerKVStore) Has(key string) (bool, error) {
	err := s.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		return err
	})
	if err == badger.ErrKeyNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *badgerKVStore) Get(key string) ([]byte, error) {
	var v []byte
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		v, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s *badgerKVStore) Set() error {
	return nil
}
