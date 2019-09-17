
package database


import (
	"github.com/jinzhu/gorm"
)


type DB interface {
	First(out interface{}, where ...interface{}) *gorm.DB
}
