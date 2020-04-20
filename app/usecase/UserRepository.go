
package usecase


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/domain"
)


type UserRepository interface {
	FindByID(id int) (event domain.Users, err error)
}
