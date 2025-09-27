//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/enkazu1116/go_home/internal/domain"
	"github.com/enkazu1116/go_home/internal/handler"
	"github.com/enkazu1116/go_home/internal/repository"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// InitializeApp はアプリケーション全体の依存関係を初期化する関数
func InitializeApp(db *gorm.DB) (*App, error) {
	wire.Build(
		// リポジトリ層の依存関係
		repository.NewTimeIsMoneyRepository,
		wire.Bind(new(repository.UserRepository), new(*repository.TimeIsMoneyGormRepo)),

		// ドメイン層の依存関係
		domain.NewUserUsecase,

		// ハンドラー層の依存関係
		handler.NewUserHandler,

		// アプリケーション全体の依存関係
		NewApp,
	)
	return &App{}, nil
}

// App はアプリケーション全体を表す構造体
type App struct {
	UserHandler *handler.UserHandler
}

// NewApp はアプリケーション全体の構造体を作成する
func NewApp(userHandler *handler.UserHandler) *App {
	return &App{
		UserHandler: userHandler,
	}
}
