package usecases

import (
	"github.com/Guiiz94/rendu-GO"
)

type ZipRepositoriesUseCase struct {
	Zipper domain.Zipper
}

func (zr *ZipRepositoriesUseCase) ZipAll(baseDir, outputPath string) error {
	return zr.Zipper.Zip(baseDir, outputPath)
}
