# Wireライブラリを使った依存性注入のサンプル

このプロジェクトでは、GoogleのWireライブラリを使用して依存性注入を実装しています。

## 実装内容

### 1. 依存関係の追加

```bash
go get github.com/google/wire/cmd/wire
go install github.com/google/wire/cmd/wire@latest
```

### 2. Wireプロバイダーの作成

`internal/wire/wire.go`で依存性注入の設定を行います：

```go
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
```

### 3. コード生成

Wireコマンドを実行して依存性注入のコードを生成します：

```bash
wire ./internal/wire
```

これにより`wire_gen.go`ファイルが生成されます。

### 4. main.goでの使用

```go
// Wireを使用した依存性注入でアプリケーションを初期化
app, err := wire.InitializeApp(db)
if err != nil {
	log.Fatalf("failed to initialize app: %v", err)
}

// HTTPサーバ設定
r := chi.NewRouter()
app.UserHandler.RegisterRoutes(r)
```

## 依存関係の流れ

1. **データベース接続** (`*gorm.DB`)
2. **リポジトリ層** (`TimeIsMoneyGormRepo` → `UserRepository`インターフェース)
3. **ドメイン層** (`UserUsecase`)
4. **ハンドラー層** (`UserHandler`)
5. **アプリケーション全体** (`App`)

## メリット

- **型安全性**: コンパイル時に依存関係のエラーを検出
- **コード生成**: 実行時にリフレクションを使用しない
- **明確な依存関係**: 依存関係がコードで明確に定義される
- **テスタビリティ**: モックやスタブの注入が容易

## 実行方法

```bash
# ビルド
go build -o app ./cmd

# 実行
./app
```

## API エンドポイント

- `POST /users` - ユーザー作成
- `GET /users` - ユーザー一覧取得
- `GET /users/{id}` - ユーザー取得
- `PUT /users/{id}` - ユーザー更新
- `DELETE /users/{id}` - ユーザー削除
