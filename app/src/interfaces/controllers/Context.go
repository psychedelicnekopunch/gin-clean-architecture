
package controllers


type Context interface {
	Cookie(name string) (string, error)
	HTML(code int, name string, obj interface{})
	JSON(code int, obj interface{})
	Param(key string) string
	Query(key string) string
	Redirect(code int, location string)
	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	ShouldBind(obj interface{}) error
}
