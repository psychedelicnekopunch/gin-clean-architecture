
package usecase


import (
	"github.com/jinzhu/gorm"

	"github.com/psychedelicnekopunch/gin-clean-architecture/src/domain"
)


type UserRepository interface {
	FindByID(db *gorm.DB, id int) (user domain.Users, err error)
}
