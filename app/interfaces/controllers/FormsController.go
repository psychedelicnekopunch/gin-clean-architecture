
package controllers


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/usecase"
)


type FormsController struct {
	Interactor usecase.FormInteractor
}


func NewFormsController() *FormsController {
	return &FormsController{
		Interactor: usecase.FormInteractor{},
	}
}


func (controller *FormsController) Get(c Context) {
	response, _ := controller.Interactor.Get()
	c.HTML(controller.Interactor.StatusCode, "forms/index.tmpl", response)
}


func (controller *FormsController) Post(c Context) {
	var params usecase.FormParameters
	c.ShouldBind(&params)
	response, _ := controller.Interactor.Post(params)

	c.HTML(controller.Interactor.StatusCode, "forms/index.tmpl", response)
}
