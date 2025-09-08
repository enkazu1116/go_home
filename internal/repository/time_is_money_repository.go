package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

type TimeIsMoneyRepository interface {
	Save(ctx context.Context, tim TimeIsMoney) error
	FindByID(ctx context.Context, id int) (*TimeIsMoney, error)
}

func (repo *timeIsMoneyGormRepo) Save(ctx context.Context, tim TimeIsMoney) error {
	return repo.db.WithContext(ctx).Save(&tim).Error
}

func (repo *timeIsMoneyGormRepo) FindByID(ctx context.Context, id int) (*TimeIsMoney, error) {
	var tim TimeIsMoney
	if err := repo.db.WithContext(ctx).First(&tim, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return &tim, nil
}

// SQLite３を使ったCRUD操作の基本確認
var Dbconnection *sql.DB

func sampleExec() {
	Dbconnection, _ := sql.Open("sqlite3", "./example.sql")
	// 遅延処理で接続解除を確実に行う
	defer Dbconnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS user(
		id STRING, 
		authId STRING, 
		name STRING, 
		email STRING,
		role STRING,
	)`

	_, err := Dbconnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}

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
