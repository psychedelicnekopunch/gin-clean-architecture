
package usecase


import (
	"errors"
)


type CookieInteractor struct {}


type CookieResponse struct {
	CookieValue string
	Forms CookieParameters
}


type CookieParameters struct {
	Value string `form:"value"`
}


func (interactor *CookieInteractor) Get(cookie string) (response CookieResponse, resultStatus *ResultStatus) {
	response.CookieValue = cookie
	return response, NewResultStatus(200, nil)
}


func (interactor *CookieInteractor) Post(cookie string, params CookieParameters) (response CookieResponse, resultStatus *ResultStatus) {
	response.CookieValue = cookie
	response.Forms = params
	if response.Forms.Value == "" {
		return response, NewResultStatus(400, errors.New("value is empty"))
	}
	response.CookieValue = response.Forms.Value
	return response, NewResultStatus(200, nil)
}
