package handler

import "go_home-main/internal/domain"

// UserHandlerの構造体を定義
// Usercaseを構造体に埋め込み、UsercaseのメソッドをUserHandlerから呼び出せるようにするs
type UserHandler struct {
	domain.UserUsecase
}

// NewUserHandlerを定義し、UserHandlerの実体を生成する
func NewUserHandler(userhandle domain.UserUsecase) *UserHandler {
	
	// 生成したUserHandlerのポインタを返す（）内は構造体のフィールド名： フィールドの値
	return &UserHandler(UserUsecase: userhandle)
}

// UserHandler用のルーティング設定
// chi.Routerはchiパッケージを呼び出し、Routerインタフェースを使用していく
func (handle *UserHandler) RegisterRoutes(router chi.Router) {
	
	// 各UserUsecaseのメソッドをHTTPメソッドとパスにマッピングする。
	// "/CreateUser"でhandler.CreateUserが呼び出される。 他も同様。
	router.Post("/CreateUser", handle.CreateUser)
}

// CreateUserのHTTPハンドラー関数を定義
// 関数の型は *UserHandlerのメソッド
// 引数はhttp.ResponseWriterと*http.Request
// 戻り値はなし
func (handler *UserHandler) CreateUser(writer http.ResponseWriter, request *http.Request) {

	// UserEntityを受け取るリクエストを定義する
	var user entity.User

	// 1. リクエストボディをデコードする（JSON -> Goの構造体）
	if err := json.NewDecoder(request.Body).Decode(&user);
}