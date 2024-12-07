package dto

import "github.com/raaaaaaaay86/ginkit/example/route/entity"

type Store struct {
	Id          int64 `json:"id"`
	TotalIncome int64 `json:"total_income"`
}

func NewStoreFromEntity(from entity.Store) *Store {
	output := Store{
		Id: from.Id,
	}

	if from.TotalIncome != nil {
		output.TotalIncome = from.TotalIncome.Load()
	}

	return &output
}

type Stores []Store

func NewStoresFromEntities(stores []entity.Store) Stores {
	if stores == nil {
		return nil
	}

	result := make(Stores, 0, len(stores))
	for _, store := range stores {
		result = append(result, *NewStoreFromEntity(store))
	}

	return result
}
