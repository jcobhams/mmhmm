package database

import (
	"sync"

	"github.com/jcobhams/mmhmm/models"
)

type (
	Database interface {
		Write(key string, value models.Note) error
		Read(key string) ([]models.Note, error)
		Delete(key string) error
		All() map[string][]models.Note
	}

	db struct {
		store map[string][]models.Note
		rm    sync.RWMutex
	}
)

func New() *db {
	return &db{
		store: make(map[string][]models.Note),
	}
}

func (d *db) Write(key string, value models.Note) error {
	d.rm.Lock()
	defer d.rm.Unlock()
	d.store[key] = append(d.store[key], value)
	return nil
}

func (d *db) Read(key string) ([]models.Note, error) {
	d.rm.RLock()
	defer d.rm.RUnlock()
	return d.store[key], nil
}

func (d *db) Delete(key string) error {
	d.rm.Lock()
	defer d.rm.Unlock()
	delete(d.store, key)
	return nil
}

// All shows the content of the database - Debugging purposes
func (d *db) All() map[string][]models.Note {
	d.rm.RLock()
	defer d.rm.RUnlock()
	return d.store
}
