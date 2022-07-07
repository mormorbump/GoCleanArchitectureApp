package domain

import "fmt"

type User struct {
	ID int
	FirstName string
	LastName string
	FullName string
}


// 仕様変更がアプリの固有の問題かどうか切り分ける。
// そしてそれはUseCaseより内側で対応するか外側で対応するかの違いとなる。
// Domain(User)に関わる仕様変更なら、アプリ固有の問題なので、domainに実装。
// これをdatabase層から直接渡してもいいが、databaseとは関係ない処理。
// なので、domainにBuild(Userの情報の整理整頓)メソッドを定義するべき。
//　こうすることでビジネスロジック(アプリ固有の実装)をusecase層より内側に閉じ込められる。

// 呼ぶところはdatabase層で良さそう。(db操作の部分がそこなので)
func (u *User) Build() *User {
	u.FullName = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return u
}

// Usersと言う名前でUser structの配列型を定義
type Users []User
