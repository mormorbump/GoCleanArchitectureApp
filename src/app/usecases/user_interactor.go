// interfaces/databaseからinput Portの役割、及び
// interfaces/controllersへのGatewayの役割を担う。
// 要するにRepositoryをラッパーしたcontrollerで実行するためのもの
package usecases

import "app/domain"

type UserInteractor struct {
	UserRepository IUserRepository
}

// Repositoryで永続化を行うため、Userを追加するメソッド
func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	id, err := interactor.UserRepository.Store(u)
	if err != nil {
		return
	}
	user, err = interactor.UserRepository.FindById(id)
	return
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}