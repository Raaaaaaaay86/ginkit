package serviceimpl

import (
	"github.com/raaaaaaaay86/ginutil/example/route/entity"
	"github.com/raaaaaaaay86/ginutil/example/route/persistence"
)

type Store struct {
	stores persistence.Store
}

func (s *Store) Create() (int64, error) {
	return s.stores.Create()
}

func (s *Store) FindAll() ([]entity.Store, error) {
	return s.stores.FindAll()
}

func (s *Store) IncrementTotalIncome(id int64) (int64, error) {
	return s.stores.IncrementTotalIncome(id)
}

func NewStore(
	stores persistence.Store,
) *Store {
	return &Store{
		stores: stores,
	}
}
