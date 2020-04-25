
package domain


import (
	"time"
)


type GithubRepositories struct {
	ID int `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}


type GithubRepositoriesForGet struct {
	ID int
	Name string
	CreatedAt time.Time
}


func (u *GithubRepositories) BuildForGet() GithubRepositoriesForGet {
	repos := GithubRepositoriesForGet{}
	repos.ID = u.ID
	repos.Name = u.Name
	repos.CreatedAt = u.CreatedAt
	return repos
}
