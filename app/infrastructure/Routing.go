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
	jsonController := controllers.NewJsonController(r.DB)
	r.Gin.GET("/json/:id", func (c *gin.Context) { jsonController.Get(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
