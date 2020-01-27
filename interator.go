package ksgo

type Iterator interface {
	Err() error
	Key() []byte
	Next() bool
	Close()
	Seek(key []byte) bool
	Value() ([]byte, error)
}
