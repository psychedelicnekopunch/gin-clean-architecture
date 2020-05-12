
package usecase


type AssetInteractor struct {
	StatusCode int
}


type AssetResponse struct {

}


func (interactor *AssetInteractor) Get() (response AssetResponse, err error) {
	interactor.StatusCode = 200
	return response, nil
}
