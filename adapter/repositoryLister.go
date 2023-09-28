package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"
	"log"

	"github.com/joho/godotenv"
	"github.com/Guiiz94/rendu-GO/domain"
)

type GithubAPIAdapter struct{}

func (g *GithubAPIAdapter) List(username string) ([]domain.Repository, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("GITHUB_TOKEN")

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s/repos", username), nil)
	if err != nil {
		return nil, err
	}

	// Si le token est disponible, on l'ajoute à la requête
	if token != "" {
		req.Header.Add("Authorization", "token "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned non-OK status: %v", resp.Status)
	}

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
