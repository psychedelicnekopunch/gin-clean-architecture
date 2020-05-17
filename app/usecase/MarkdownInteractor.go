
package usecase


type MarkdownInteractor struct {
	StatusCode int
}


type MarkdownResponse struct {
	Value string
}


func (interactor *MarkdownInteractor) Get() (response MarkdownResponse, err error) {
	response.Value = `
## title

* list1
* list2


### link

https://github.com/psychedelicnekopunch/gin-clean-architecture
`
	interactor.StatusCode = 200
	return response, nil
}
