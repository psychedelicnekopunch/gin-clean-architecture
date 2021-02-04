
package usecase


type MarkdownInteractor struct {}


type MarkdownResponse struct {
	Value string
}


func (interactor *MarkdownInteractor) Get() (response MarkdownResponse, resultStatus *ResultStatus) {
	response.Value = `
## title

* list1
* list2


### link

https://github.com/psychedelicnekopunch/gin-clean-architecture
`
	return response, NewResultStatus(200, nil)
}
