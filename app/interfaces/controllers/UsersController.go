
package controllers


import (
	"strconv"

	"psychedelicnekopunch/gin-clean-architecture/api/app/interfaces/database"
	"psychedelicnekopunch/gin-clean-architecture/api/app/usecase"
)


type UsersController struct {
	Interactor usecase.UserInteractor
}


func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor{
			User: &database.UserRepository{ DB: db },
		},
	}
}


func (controller *UsersController) Get(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.Interactor.Get(id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", user))
}
