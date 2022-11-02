package store

import (
	"fmt"
)

type RepositoryInMem struct {
	m map[string]*Store
}

func NewInMemRepository() Repository {
	return &RepositoryInMem{
		m: map[string]*Store{},
	}
}

func (w *RepositoryInMem) AddStore(Store Store) error {
	if _, ok := w.m[Store.Subdomain]; !ok {
		w.m[Store.Subdomain] = &Store
		return nil
	} else {
		return fmt.Errorf("store %s already added", Store.Subdomain)
	}
}

func (w *RepositoryInMem) GetStore(subdomain string) (Store, error) {
	if s, ok := w.m[subdomain]; ok {
		return *s, nil
	} else {
		return Store{}, fmt.Errorf("websiteserver not found")
	}
}

func (w *RepositoryInMem) UpdateStore() error {
	//TODO implement me
	panic("implement me")
}

func (w *RepositoryInMem) DeleteStore(subdomain string) error {
	delete(w.m, subdomain)
	return nil
}
