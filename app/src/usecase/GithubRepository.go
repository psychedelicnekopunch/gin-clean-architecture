
package usecase


import (
	"github.com/psychedelicnekopunch/gin-clean-architecture/src/domain"
)


type GithubRepository interface {
	GetRepositoryList(params map[string]string) (repos []domain.GithubRepositories, err error)
}
