package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Guiiz94/rendu-GO/domain"
)

type GithubAPIAdapter struct{}

func (g *GithubAPIAdapter) List(username string) ([]domain.Repository, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiRepos []struct {
		Name       string    `json:"name"`
		HTMLURL    string    `json:"html_url"`
		UpdatedAt  time.Time `json:"updated_at"`
		CloneURL   string    `json:"clone_url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&apiRepos); err != nil {
		return nil, err
	}

	var repos []domain.Repository
	for _, apiRepo := range apiRepos {
		repos = append(repos, domain.Repository{
			Name:       apiRepo.Name,
			HTMLURL:    apiRepo.HTMLURL,
			UpdatedAt: apiRepo.UpdatedAt.Format(time.RFC3339),
			CloneURL:   apiRepo.CloneURL,
		})
	}

	return repos, nil
}
