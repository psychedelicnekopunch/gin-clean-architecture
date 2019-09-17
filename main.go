package main

import (
	"psychedelicnekopunch/gin-clean-architecture/api/app/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run(":8080")
}
