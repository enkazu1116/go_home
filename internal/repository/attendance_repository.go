package repository

import (
	"context"
	"errors"

	"github.com/enkazu1116/go_home/internal/entity"
	"gorm.io/gorm"
)

// AttendanceRepository は勤怠エンティティのリポジトリインターフェース
type AttendanceRepository interface {
	Create(ctx context.Context, a entity.Attendance) error
	Update(ctx context.Context, a entity.Attendance) error
	FindByID(ctx context.Context, id string) (*entity.Attendance, error)
	FindByUserID(ctx context.Context, userID string) ([]entity.Attendance, error)
	FindAll(ctx context.Context) ([]entity.Attendance, error)
	Delete(ctx context.Context, a entity.Attendance) error
}

// Gorm実装
type attendanceGormRepo struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceGormRepo{db: db}
}

func (r *attendanceGormRepo) Create(ctx context.Context, a entity.Attendance) error {
	return r.db.WithContext(ctx).Create(&a).Error
}

func (r *attendanceGormRepo) Update(ctx context.Context, a entity.Attendance) error {
	return r.db.WithContext(ctx).Save(&a).Error
}

func (r *attendanceGormRepo) FindByID(ctx context.Context, id string) (*entity.Attendance, error) {
	var a entity.Attendance
	err := r.db.WithContext(ctx).First(&a, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("attendance not found")
		}
		return nil, err
	}
	return &a, nil
}

func (r *attendanceGormRepo) FindByUserID(ctx context.Context, userID string) ([]entity.Attendance, error) {
	var list []entity.Attendance
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (r *attendanceGormRepo) FindAll(ctx context.Context) ([]entity.Attendance, error) {
	var list []entity.Attendance
	err := r.db.WithContext(ctx).Find(&list).Error
	return list, err
}

func (r *attendanceGormRepo) Delete(ctx context.Context, a entity.Attendance) error {
	return r.db.WithContext(ctx).Delete(&a).Error
}
