
package controllers


import (
	"strconv"

	"github.com/psychedelicnekopunch/gin-clean-architecture/src/interfaces/database"
	"github.com/psychedelicnekopunch/gin-clean-architecture/src/usecase"
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

	user, res := controller.Interactor.Get(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", user))
}
