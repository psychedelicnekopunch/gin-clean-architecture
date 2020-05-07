
package controllers


import (
	"fmt"
)


type IndexController struct {}


func NewIndexController() *IndexController {
	return &IndexController{}
}


func (controller *IndexController) Get(c Context) {
	fmt.Print(c.Query("test"))
	c.HTML(200, "index.tmpl", nil)
}
