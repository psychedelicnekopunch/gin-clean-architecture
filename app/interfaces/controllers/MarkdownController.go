
package controllers


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/usecase"
)


type MarkdownController struct {
	Interactor usecase.MarkdownInteractor
}


func NewMarkdownController() *MarkdownController {
	return &MarkdownController{
		Interactor: usecase.MarkdownInteractor{},
	}
}


func (controller *MarkdownController) Get(c Context) {
	response, _ := controller.Interactor.Get()
	c.HTML(controller.Interactor.StatusCode, "markdown/index.tmpl", response)
}
