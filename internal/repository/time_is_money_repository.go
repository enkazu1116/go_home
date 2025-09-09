package repository

import (
	"context"
	"errors"

	"go_home-main/internal/entity"

	"gorm.io/gorm"
)

// Gormの実装練習
// コメントは過剰につけるが、実際のコードでは適さない。
// 理解のため、何が引数で、戻り値は何かを一目でわかるようにしている。

// リポジトリの構造体
type timeIsMoneyGormRepo struct {
	// Gormを用いたDB接続
	db *gorm.DB
}

// コンストラクタ
// DB接続を受け取る
// 引数： db(*gorm.DB型)
// 戻り値: *timeIsMoneyGormRepo型のポインタ
func NewTimeIsMoneyRepository(db *gorm.DB) *timeIsMoneyGormRepo {
	return &timeIsMoneyGormRepo{db: db}
}

// Gormを用いたCRUD操作のメソッドを定義
// 新規登録 INSERT
// 引数: user(User型)
// 戻り値: error型
func (repo *timeIsMoneyGormRepo) CreateUser(context context.Context, user entity.User) error {
	// エラーハンドリングは呼び出し元で行う
	// Createは新規登録のため、アドレスを渡す。
	return repo.db.WithContext(context).Create(&user).Error
}

// 更新処理 UPDATE
// 引数: user(User型)
// 戻り値: error型
func (repo *timeIsMoneyGormRepo) UpdateUser(context context.Context, user entity.User) error {
	// エラーハンドリングは呼び出し元で行う
	// Saveは主キーを元にレコードを更新するため、アドレスを渡す
	return repo.db.WithContext(context).Save(&user).Error
}

// 取得処理 SELECT (1件)
// 引数: context(context.Context型), user(User型)
// 戻り値: (*User, error) ※ここでerrorだけを返すと、userを返すことができないため、*Userも返す
func (repo *timeIsMoneyGormRepo) FindFirst(context context.Context, id string) (*entity.User, error) {

	// 検索結果を格納する
	var user entity.User

	// 失敗 errorを変数名にすると、Go言語の組み込み型と衝突する
	// IDを元に検索する
	// Firstは最初の1件を取得する
	err := repo.db.WithContext(context).First(&user, "id = ?", id).Error

	// 条件に一致しない場合は、errにエラーが格納される
	if err != nil {
		// レコードが見つからない場合のエラーハンドリング
		// ErrRecordNotFoundはGormが提供するエラー
		// errorsは標準パッケージで、Is関数を使ってエラーの種類を判定する
		if errors.Is(err, gorm.ErrRecordNotFound) {

			// New関数でカスタムエラーメッセージを作成して返す。
			return nil, errors.New("検索結果が見つかりません。")
		}
	}

	// 正常に取得できた場合は、userとnil(err)を返す
	return &user, nil
}

// 取得処理 SELECT (複数)
// 引数: context(context.Context型), user(User型)
// 戻り値: (*User, error)
func (repo *timeIsMoneyGormRepo) FindAllUser(context context.Context) ([]entity.User, error) {

	// 検索結果を格納するスライス
	var users []entity.User

	// Findは条件に一致する全てのレコードを取得する
	err := repo.db.WithContext(context).Find(&users).Error
	if err != nil {
		// レコードが見つからない場合のエラーハンドリング
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ユーザーが登録されていません。")
		}
	}
	return users, nil
}

// 削除処理 DELETE
// 引数: context(context.Context型), user(User型)
// 戻り値: (error)
func (repo *timeIsMoneyGormRepo) DeleteUser(context context.Context, user entity.User) error {
	// Deleteはレコードを削除する
	return repo.db.WithContext(context).Delete(&user).Error
}
