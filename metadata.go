package hoverfly

import (
	"bytes"
	"fmt"

	"github.com/boltdb/bolt"
)

// Metadata - interface to store and retrieve any metadata that is related to Hoverfly
type Metadata interface {
	Set(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	GetAll() ([]MetaObject, error)
	CloseDB()
}

