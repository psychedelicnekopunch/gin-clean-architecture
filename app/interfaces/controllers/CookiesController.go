
package controllers


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/usecase"
)


type CookiesController struct {
	Interactor usecase.CookieInteractor
}


func NewCookiesController() *CookiesController {
	return &CookiesController{
		Interactor: usecase.CookieInteractor{},
	}
}


func (controller *CookiesController) Get(c Context) {
	cookie, _ := c.Cookie("value")
	response, res := controller.Interactor.Get(cookie)
	c.HTML(res.StatusCode, "cookies/index.tmpl", response)
}


func (controller *CookiesController) Post(c Context) {
	cookie, _ := c.Cookie("value")
	var params usecase.CookieParameters
	c.ShouldBind(&params)
	response, res := controller.Interactor.Post(cookie, params)

	if res.Error == nil {
		c.SetCookie("value", response.CookieValue, 60*60*24*365, "/", "", false, true)
	}

	c.HTML(res.StatusCode, "cookies/index.tmpl", response)
}
