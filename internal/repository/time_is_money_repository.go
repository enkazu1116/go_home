package repository

import (
	"context"
	"errors"
)

// TimeIsMoney エンティティの例
type TimeIsMoney struct {
	ID    int
	Value int
}

// TimeIsMoneyRepository インターフェース
type TimeIsMoneyRepository interface {
	Save(ctx context.Context, tim TimeIsMoney) error
	FindByID(ctx context.Context, id int) (*TimeIsMoney, error)
}

// インメモリ実装例
type timeIsMoneyRepoImpl struct {
	store map[int]TimeIsMoney
}

func NewTimeIsMoneyRepository() TimeIsMoneyRepository {
	return &timeIsMoneyRepoImpl{
		store: make(map[int]TimeIsMoney),
	}
}

func (r *timeIsMoneyRepoImpl) Save(ctx context.Context, tim TimeIsMoney) error {
	r.store[tim.ID] = tim
	return nil
}

func (r *timeIsMoneyRepoImpl) FindByID(ctx context.Context, id int) (*TimeIsMoney, error) {
	tim, ok := r.store[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &tim, nil
}
