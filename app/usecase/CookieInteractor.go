
package usecase


import (
	"errors"
)


type CookieInteractor struct {
	StatusCode int
}


type CookieResponse struct {
	CookieValue string
	Forms CookieParameters
}


type CookieParameters struct {
	Value string `form:"value"`
}


func (interactor *CookieInteractor) Get(cookie string) (response CookieResponse, err error) {
	response.CookieValue = cookie
	interactor.StatusCode = 200
	return response, nil
}


func (interactor *CookieInteractor) Post(cookie string, params CookieParameters) (response CookieResponse, err error) {
	response.CookieValue = cookie
	response.Forms = params
	if response.Forms.Value == "" {
		interactor.StatusCode = 400
		return response, errors.New("value is empty")
	}
	response.CookieValue = response.Forms.Value
	interactor.StatusCode = 200
	return response, nil
}
