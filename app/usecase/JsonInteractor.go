
package usecase


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/domain"
)


type JsonInteractor struct {
	DB DBRepository
	User UserRepository
	StatusCode int
}


func (interactor *JsonInteractor) Get(id int) (user domain.UsersForGet, err error) {
	db := interactor.DB.Connect()
	// Users の取得
	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		interactor.StatusCode = 404
		return domain.UsersForGet{}, err
	}
	user = foundUser.BuildForGet()
	interactor.StatusCode = 200
	return user, nil
}
