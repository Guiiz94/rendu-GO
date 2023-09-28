package adapters

import (
	"encoding/csv"
	"os"

	"github.com/rendu-GO/domain"
)

type CSVAdapter struct{}

func (c *CSVAdapter) Write(repos []domain.Repository) error {
	file, err := os.Create("repositories.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, repo := range repos {
		err := writer.Write([]string{repo.Name, repo.HTMLURL, repo.UpdatedAt.String(), repo.CloneURL})
		if err != nil {
			return err
		}
	}

	return nil
}