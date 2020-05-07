
package controllers


type Context interface {
	Query(key string) string
	Param(key string) string
	JSON(code int, obj interface{})
	HTML(code int, name string, obj interface{})
	ShouldBind(obj interface{}) error
	Cookie(name string) (string, error)
	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
}
