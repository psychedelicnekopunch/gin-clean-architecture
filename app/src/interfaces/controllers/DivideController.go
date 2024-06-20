
package controllers


import (
	"strconv"
	"github.com/gin-gonic/gin"
)


type DividesController struct {}


func NewDividesController() *DividesController {
	return new(DividesController)
}


func (controller *DividesController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, NewH("success", 10 / id))
}
