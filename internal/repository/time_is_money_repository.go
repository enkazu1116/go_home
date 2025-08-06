package repository

import (
	"context"
	"time"

	"immediately_return_home/internal/entity"
)

// Repositoryパターンのインターフェース定義
type timeIsMoneyRepository interface {
	// 出勤記録を保存
	SaveCheckIn(ctx context.Context, userID string, checkIn time.Time) error

	// 退勤記録を保存
	SaveCheckOut(ctx context.Context, userID string, checkOut time.Time) error

	// ユーザーIDと年月で勤怠データを取得
	GetMonthlyAttendance(ctx context.Context, userID string, year int, month int) ([]*entity.Attendance, error)

	// 勤怠情報を更新
	UpdateAttendance(ctx context.Context, attendance *entity.Attendance) error
}
