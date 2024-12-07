package persistenceimpl

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/raaaaaaaay86/ginkit/example/route/entity"
)

type Store struct {
	currentId *atomic.Int64
	stores    *sync.Map
}

func (s *Store) Create() (int64, error) {
	id := s.currentId.Add(1)

	store := entity.Store{
		Id:          id,
		TotalIncome: &atomic.Int64{},
	}

	s.stores.Store(id, &store)

	return id, nil
}

func (s *Store) FindAll() ([]entity.Store, error) {
	var err error

	stores := make([]entity.Store, 0)

	s.stores.Range(func(key, value any) bool {
		store, ok := value.(*entity.Store)
		if !ok {
			err = fmt.Errorf("500|unexpected value type: %T", value)
			return false
		}

		stores = append(stores, *store)

		return true
	})

	if err != nil {
		return nil, err
	}

	return stores, nil
}

func (s *Store) IncrementTotalIncome(id int64) (int64, error) {
	value, ok := s.stores.Load(id)
	if !ok {
		return 0, fmt.Errorf("404|store not found: %d", id)
	}

	store, ok := value.(*entity.Store)
	if !ok {
		return 0, fmt.Errorf("500|unexpected value type: %T", value)
	}

	if store.TotalIncome == nil {
		store.TotalIncome = &atomic.Int64{}
	}

	return store.TotalIncome.Add(1), nil
}

func NewStore() *Store {
	return &Store{
		currentId: &atomic.Int64{},
		stores:    &sync.Map{},
	}
}
