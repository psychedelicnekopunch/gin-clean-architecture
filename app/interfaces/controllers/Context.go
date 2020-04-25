
package controllers


type Context interface {
	Param(key string) string
	JSON(code int, obj interface{})
	HTML(code int, name string, obj interface{})
	ShouldBind(obj interface{}) error
}
