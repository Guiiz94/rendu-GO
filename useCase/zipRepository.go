package usecases

import (
	"rendu-GO/domain"
)

type ZipRepositoriesUseCase struct {
	Zipper domain.Zipper
}

func (zr *ZipRepositoriesUseCase) ZipAll(baseDir, outputPath string) error {
	return zr.Zipper.Zip(baseDir, outputPath)
}
