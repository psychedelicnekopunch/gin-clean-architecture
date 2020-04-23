
package usecase


import (
	"github.com/jinzhu/gorm"

	"github.com/psychedelicnekopunch/gin-clean-architecture/app/domain"
)


type UserRepository interface {
	FindByID(db *gorm.DB, id int) (event domain.Users, err error)
}
