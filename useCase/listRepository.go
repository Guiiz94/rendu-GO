package usecases

import (
	"github.com/Guiiz94/rendu-GO/domain"
	"sort"
	"time"
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
		timeI, errI := time.Parse(time.RFC3339, repos[i].UpdatedAt)
		timeJ, errJ := time.Parse(time.RFC3339, repos[j].UpdatedAt)
		
		// Gérer les erreurs de parsing si nécessaire
		if errI != nil || errJ != nil {
			return false  // ou un autre comportement par défaut
		}
	
		return timeI.After(timeJ)
	})
	

	if len(repos) > limit {
		return repos[:limit], nil
	}
	return repos, nil
}
