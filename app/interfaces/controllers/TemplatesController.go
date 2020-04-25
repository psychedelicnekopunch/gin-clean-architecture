
package controllers


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/interfaces/database"
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/usecase"
)


type TemplatesController struct {
	Interactor usecase.TemplateInteractor
}


func NewTemplatesController(http database.Http) *TemplatesController {
	return &TemplatesController{
		Interactor: usecase.TemplateInteractor{
			Github: &database.GithubRepository{ Http: http },
		},
	}
}


func (controller *TemplatesController) Get(c Context) {
	response, err := controller.Interactor.Get()
	if err != nil {
		c.HTML(controller.Interactor.StatusCode, "templates/error.tmpl", response)
		return
	}
	c.HTML(controller.Interactor.StatusCode, "templates/index.tmpl", response)
}
