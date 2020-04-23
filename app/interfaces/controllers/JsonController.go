
package controllers


import (
	"strconv"

	"github.com/psychedelicnekopunch/gin-clean-architecture/app/interfaces/database"
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/usecase"
)


type JsonController struct {
	Interactor usecase.JsonInteractor
}


func NewJsonController(db database.DB) *JsonController {
	return &JsonController{
		Interactor: usecase.JsonInteractor{
			DB: &database.DBRepository{ DB: db },
			User: &database.UserRepository{},
		},
	}
}


func (controller *JsonController) Get(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.Interactor.Get(id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH(err.Error(), nil))
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH("success", user))
}
