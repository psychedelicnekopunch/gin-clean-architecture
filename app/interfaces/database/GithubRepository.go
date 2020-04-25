
package database


import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/psychedelicnekopunch/gin-clean-architecture/app/domain"
)


type GithubRepository struct {
	Http Http
}


func (repo *GithubRepository) GetRepositoryList(params map[string]string) (repos []domain.GithubRepositories, err error) {

	request, err := repo.Http.Get("https://api.github.com/users/psychedelicnekopunch/repos", params)
	response, err := repo.Http.GetResponse(request)

	if err != nil {
		return []domain.GithubRepositories{}, err
	}

	body, err := repo.Http.GetBody(response)

	if err != nil {
		return []domain.GithubRepositories{}, err
	}

	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&repos)

	if len(repos) == 0 {
		return repos, errors.New("not found repos")
	}

	return repos, nil
}
