
package usecase


type FormInteractor struct {}


type FormResponse struct {
	Method string
	Value string
	Forms FormParameters
}


type FormParameters struct {
	Value string `form:"value"`
}


func (interactor *FormInteractor) Get() (response FormResponse, resultStatus *ResultStatus) {
	response.Method = "GET"
	return response, NewResultStatus(200, nil)
}

func (interactor *FormInteractor) Post(params FormParameters) (response FormResponse, resultStatus *ResultStatus) {
	response.Method = "POST"
	response.Forms = params
	return response, NewResultStatus(200, nil)
}
