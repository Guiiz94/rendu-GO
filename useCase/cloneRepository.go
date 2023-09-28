package usecases

import (
	"fmt"

	"github.com/Guiiz94/rendu-GO/domain"
)

type CloneRepositoriesUseCase struct {
	Git domain.GitManager
}

func (cr *CloneRepositoriesUseCase) CloneAll(repos []domain.Repository, baseDir string) error {
	for _, repo := range repos {
		dir := baseDir + "/" + repo.Name
		if err := cr.Git.Clone(repo, dir); err != nil {
			fmt.Println(repo)
			fmt.Println(dir)
			return err
		}
	}
	return nil
}
