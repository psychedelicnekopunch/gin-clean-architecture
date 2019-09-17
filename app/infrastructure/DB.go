
package infrastructure


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type DB struct {
	Host string
	Username string
	Password string
	DBName string
	Connect *gorm.DB
}


func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host: c.DB.Production.Host,
		Username: c.DB.Production.Username,
		Password: c.DB.Production.Password,
		DBName: c.DB.Production.DBName,
	})
}


func NewTestDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host: c.DB.Test.Host,
		Username: c.DB.Test.Username,
		Password: c.DB.Test.Password,
		DBName: c.DB.Test.DBName,
	})
}


func newDB(d *DB) *DB {
	// https://github.com/go-sql-driver/mysql#examples
	db, err := gorm.Open("mysql", d.Username + ":" + d.Password + "@tcp(" + d.Host + ")/" + d.DBName + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	d.Connect = db
	return d
}


func (db *DB) First(out interface{}, where ...interface{}) *gorm.DB {
	return db.Connect.First(out, where...)
}
