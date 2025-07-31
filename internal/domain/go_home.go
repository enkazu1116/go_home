package domain

import (
	"context"
	"time"
)

// 出勤・退勤APIのインターフェース
type AttendanceUsecase interface {

	// 出勤・退勤処理
	CheckIn(ctx context.Context, userID string, checkInTime time.Time) error
	CheckOut(ctx context.Context, userID string, checkOutTime time.Time) error

	// 月毎の情報を取得・変更処理
	GetMontlyAttendance(ctx context.Context, userID string, enterDate time.Time) error
	UpdateMonthlyAttendance(ctx context.Context, userID string, enterDate time.Time) error
}
