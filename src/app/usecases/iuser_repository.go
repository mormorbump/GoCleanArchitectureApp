package usecases

import "app/domain"

//DIP使ってリポジトリの内容を定義
// interfaces/database/user_repositoryからのインプットをここで実現し、
// interfaces/controllersへのGatewayをusercase/user_interactor.goで実現
type IUserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}