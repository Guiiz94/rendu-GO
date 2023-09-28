package usecases

import (
	"github.com/rendu-GO/domain"
)

type UpdateRepositoriesUseCase struct {
	Git domain.GitManager
}

func (ur *UpdateRepositoriesUseCase) UpdateAll(repos []domain.Repository, baseDir string) error {
	for _, repo := range repos {
		dir := baseDir + "/" + repo.Name
		if err := ur.Git.Update(dir); err != nil {
			return err
		}
	}
	return nil
}
