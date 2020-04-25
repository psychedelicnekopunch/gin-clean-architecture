
package controllers


type IndexController struct {}


func NewIndexController() *IndexController {
	return &IndexController{}
}


func (controller *IndexController) Get(c Context) {
	c.HTML(200, "index.tmpl", nil)
}
