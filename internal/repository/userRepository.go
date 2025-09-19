package repository

import (
	"context"

	"github.com/enkazu1116/go_home/internal/entity"
)

type UserRepository interface {
	// 新規登録
	CreateUser(ctx context.Context, user entity.User) error

	// 更新
	UpdateUser(ctx context.Context, user entity.User) error

	// 最初の1件を取得
	FindFirst(ctx context.Context, id string) (*entity.User, error)

	// 全件取得
	FindAllUser(ctx context.Context) ([]entity.User, error)

	// 削除
	DeleteUser(ctx context.Context, user entity.User) error
}
