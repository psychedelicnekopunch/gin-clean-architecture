package main

import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run()
}
