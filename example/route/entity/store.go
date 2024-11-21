package entity

import "sync/atomic"

type Store struct {
	Id          int64         `json:"id"`
	TotalIncome *atomic.Int64 `json:"total_income"`
}
