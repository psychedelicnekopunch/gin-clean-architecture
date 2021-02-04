
package controllers


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/src/usecase"
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
	response, res := controller.Interactor.Get()
	c.HTML(res.StatusCode, "markdown/index.tmpl", response)
}
