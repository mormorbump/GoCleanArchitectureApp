package database

import (
	"app/domain"

)

type UserRepository struct {
	ISqlHandler // infrastructureのSqlHandlerもISqlHandlerインターフェースを実装しているので、このインターフェースを依存させるだけでSqlHandlerのメソッドがつかえる。
}

// 保存処理
func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INT users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	defer row.Close() // 最後に行を閉じる
	if err != nil {
		return
	}

	var id int
	var firstName string
	var lastName string
	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return
	}

	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	user.Build() // domain層の処理
	return
}

// usersとか返り値のとこで変数定義しちゃってる
func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, first_name, last_name FROM users")
	defer rows.Close() // 最後に行を閉じる
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := domain.User{
			ID: id,
			FirstName: firstName,
			LastName: lastName,
		}
		user.Build() // domain層の処理
		users = append(users, user)
	}
	return
}

