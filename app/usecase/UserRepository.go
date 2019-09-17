
package usecase


import (
	"psychedelicnekopunch/gin-clean-architecture/api/app/domain"
)


type UserRepository interface {
	FindByID(id int) (event domain.Users, err error)
}
