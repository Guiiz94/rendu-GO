package usecases

import (
	"github.com/rendu-GO/domain"
	"sort"
)

type ListRepositoriesUseCase struct {
	RepoLister domain.RepositoryLister
}

func (lr *ListRepositoriesUseCase) List(username string, limit int) ([]domain.Repository, error) {
	repos, err := lr.RepoLister.List(username)
	if err != nil {
		return nil, err
	}

	sort.Slice(repos, func(i, j int) bool {
		return repos[i].UpdatedAt.After(repos[j].UpdatedAt)
	})

	if len(repos) > limit {
		return repos[:limit], nil
	}
	return repos, nil
}
