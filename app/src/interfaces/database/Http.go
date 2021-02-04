
package database


import (
	"net/http"
)


type Http interface {
	GetResponse(request *http.Request) (response *http.Response, err error)
	GetBody(response *http.Response) (body string, err error)
	Get(URL string, params map[string]string) (request *http.Request, err error)
	Post(URL string, params map[string]string) (request *http.Request, err error)
	Delete(URL string, params map[string]string) (request *http.Request, err error)
}
