
package usecase


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/src/domain"
)


type JsonInteractor struct {
	DB DBRepository
	User UserRepository
}


func (interactor *JsonInteractor) Get(id int) (user domain.UsersForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// Users の取得
	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		return domain.UsersForGet{},NewResultStatus(404, err)
	}
	user = foundUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}
