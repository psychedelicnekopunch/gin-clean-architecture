
package usecase


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/app/domain"
)


type TemplateInteractor struct {
	Github GithubRepository
	StatusCode int
}


type TemplateResponse struct {
	ErrorMessage string
	Lists []domain.GithubRepositoriesForGet
}


func (interactor *TemplateInteractor) Get() (response TemplateResponse, err error) {

	params := map[string]string{}
	params["sort"] = "created"
	params["direction"] = "desc"

	repos, err := interactor.Github.GetRepositoryList(params)

	if err != nil {
		response.ErrorMessage = err.Error()
		interactor.StatusCode = 400
		return response, err
	}

	for _, repo := range repos {
		response.Lists = append(response.Lists, repo.BuildForGet())
	}

	interactor.StatusCode = 200
	return response, nil
}
