
package controllers


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/src/interfaces/database"
	"github.com/psychedelicnekopunch/gin-clean-architecture/src/usecase"
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
	response, res := controller.Interactor.Get()
	if res.Error != nil {
		c.HTML(res.StatusCode, "templates/error.tmpl", response)
		return
	}
	c.HTML(res.StatusCode, "templates/index.tmpl", response)
}
