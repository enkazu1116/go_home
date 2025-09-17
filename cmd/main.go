package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go_home-main/internal/domain"
	"go_home-main/internal/entity"
	"go_home-main/internal/handler"
	"go_home-main/internal/repository"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// DB初期化（SQLiteを使用）
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	// マイグレーション
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	// リポジトリ・ユースケース・ハンドラ初期化
	repo := repository.NewTimeIsMoneyRepository(db) // concrete impl that satisfies repository.UserRepository
	usecase := domain.NewUserUsecase(repo)
	userHandler := handler.NewUserHandler(usecase)

	// HTTPサーバ設定
	r := chi.NewRouter()
	userHandler.RegisterRoutes(r)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// graceful shutdown 準備
	idleConnsClosed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		// 5秒以内にクリーンにシャットダウン
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Printf("starting server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	log.Println("server stopped")
}
