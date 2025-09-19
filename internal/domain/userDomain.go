package domain

import (
	"context"

	"github.com/enkazu1116/go_home/internal/entity"
	"github.com/enkazu1116/go_home/internal/repository"
)

// Userを使用してお試し
// ユーザーユースケースのインターフェースを定義
type UserUsecase interface {

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

// ユーザーユースケースの構造体を定義
type userUsecase struct {
	repo repository.UserRepository
}

// 新規登録呼び出し
func (u *userUsecase) CreateUser(ctx context.Context, user entity.User) error {
	return u.repo.CreateUser(ctx, user)
}

// 削除処理呼び出し
func (u *userUsecase) DeleteUser(ctx context.Context, user entity.User) error {
	return u.repo.DeleteUser(ctx, user)
}

// 全件取得呼び出し
func (u *userUsecase) FindAllUser(ctx context.Context) ([]entity.User, error) {
	return u.repo.FindAllUser(ctx)
}

// 最初の1件取得呼び出し
func (u *userUsecase) FindFirst(ctx context.Context, id string) (*entity.User, error) {
	return u.repo.FindFirst(ctx, id)
}

// 更新処理呼び出し
func (u *userUsecase) UpdateUser(ctx context.Context, user entity.User) error {
	return u.repo.UpdateUser(ctx, user)
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}
