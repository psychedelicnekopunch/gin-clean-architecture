package main

import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	http := infrastructure.NewHttp()
	r := infrastructure.NewRouting(db, http)
	r.Run()
}
