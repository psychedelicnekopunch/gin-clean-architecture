
package controllers


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/src/usecase"
)


type AssetsController struct {
	Interactor usecase.AssetInteractor
}


func NewAssetsController() *AssetsController {
	return &AssetsController{
		Interactor: usecase.AssetInteractor{},
	}
}


func (controller *AssetsController) Get(c Context) {
	response, res := controller.Interactor.Get()
	c.HTML(res.StatusCode, "assets/index.tmpl", response)
}
