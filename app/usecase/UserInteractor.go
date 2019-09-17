
package usecase


import (
	"psychedelicnekopunch/gin-clean-architecture/api/app/domain"
)


type UserInteractor struct {
	User UserRepository
	StatusCode int
}


func (interactor *UserInteractor) Get(id int) (user domain.UsersForGet, err error) {
	// Users の取得
	foundUser, err := interactor.User.FindByID(id)
	if err != nil {
		interactor.StatusCode = 404
		return domain.UsersForGet{}, err
	}
	user = foundUser.BuildForGet()
	interactor.StatusCode = 200
	return user, nil
}
