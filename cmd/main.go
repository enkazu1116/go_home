package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/enkazu1116/go_home/internal/entity"
	"github.com/enkazu1116/go_home/internal/wire"

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

	// Wireを使用した依存性注入でアプリケーションを初期化
	app, err := wire.InitializeApp(db)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	// HTTPサーバ設定
	r := chi.NewRouter()
	app.UserHandler.RegisterRoutes(r)

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
