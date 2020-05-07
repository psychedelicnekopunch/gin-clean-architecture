
package usecase


type FormInteractor struct {
	StatusCode int
}


type FormResponse struct {
	Method string
	Value string
	Forms FormParameters
}


type FormParameters struct {
	Value string `form:"value"`
}


func (interactor *FormInteractor) Get() (response FormResponse, err error) {
	response.Method = "GET"
	interactor.StatusCode = 200
	return response, nil
}

func (interactor *FormInteractor) Post(params FormParameters) (response FormResponse, err error) {
	response.Method = "POST"
	response.Forms = params
	interactor.StatusCode = 200
	return response, nil
}