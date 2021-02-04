
package usecase


type AssetInteractor struct {}


type AssetResponse struct {}


func (interactor *AssetInteractor) Get() (response AssetResponse, resultStatus *ResultStatus) {
	return response, NewResultStatus(200, nil)
}
