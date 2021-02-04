
package usecase


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/domain"
)


type TemplateInteractor struct {
	Github GithubRepository
}


type TemplateResponse struct {
	ErrorMessage string
	Lists []domain.GithubRepositoriesForGet
}


func (interactor *TemplateInteractor) Get() (response TemplateResponse, resultStatus *ResultStatus) {

	params := map[string]string{}
	params["sort"] = "created"
	params["direction"] = "desc"

	repos, err := interactor.Github.GetRepositoryList(params)

	if err != nil {
		response.ErrorMessage = err.Error()
		return response, NewResultStatus(400, err)
	}

	for _, repo := range repos {
		response.Lists = append(response.Lists, repo.BuildForGet())
	}

	return response, NewResultStatus(200, nil)
}
