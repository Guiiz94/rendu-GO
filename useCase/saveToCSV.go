package usecases

import (
	"github.com/Guiiz94/rendu-GO/domain"
)

type SaveToCSVUseCase struct {
	Writer domain.CSVWriter
}

func (sc *SaveToCSVUseCase) Save(repos []domain.Repository) error {
	return sc.Writer.Write(repos)
}
