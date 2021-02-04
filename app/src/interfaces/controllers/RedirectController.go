
package controllers


type RedirectController struct {}


func NewRedirectController() *RedirectController {
	return &RedirectController{}
}


func (controller *RedirectController) Get(c Context) {
	// Redirect は 30x系じゃないとエラーが起きる？
	// c.Redirect(201, "/")
	// c.Redirect(301, "/")
	// c.Redirect(302, "/")
	c.Redirect(303, "/")
}
