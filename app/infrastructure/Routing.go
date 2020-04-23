package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/psychedelicnekopunch/gin-clean-architecture/app/interfaces/controllers"
)

type Routing struct {
	DB *DB
	Gin *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	c := NewConfig()
	r := &Routing{
		DB: db,
		Gin: gin.Default(),
		Port: c.Routing.Port,
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	usersController := controllers.NewUsersController(r.DB)
	r.Gin.GET("/users/:id", func (c *gin.Context) { usersController.Get(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
