
package controllers


import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)


type DividesController struct {}


func NewDividesController() *DividesController {
	return new(DividesController)
}


func (controller *DividesController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var i float64 = 10
	res := i / float64(id)
	// var i int = 10
	// res := i / id
	c.JSON(200, NewH("success", fmt.Sprintf("%G / %G = %G", i, float64(id), res)))
}
