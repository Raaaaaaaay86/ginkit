package service

import "github.com/raaaaaaaay86/ginkit/example/route/entity"

type Store interface {
	Create() (int64, error)
	FindAll() ([]entity.Store, error)
	IncrementTotalIncome(id int64) (int64, error)
}