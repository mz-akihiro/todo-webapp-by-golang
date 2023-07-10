package model

// トークン認証用ミドルウェアとその後に実行されるハンドラーにuser idを渡すための型
type ContextKey struct {
	UserId string
}
