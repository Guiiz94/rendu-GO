package adapters

import (
	"fmt"
	"os/exec"

	"github.com/Guiiz94/rendu-GO/domain"
)

type GitAdapter struct{}

func (g *GitAdapter) Clone(repo domain.Repository, dir string) error {
	cmd := exec.Command("git", "clone", repo.CloneURL, dir)
	fmt.Printf("Cloning repo: %s. Directory: %s\n", repo.Name, dir)
	fmt.Printf("Command: %s\n", cmd.String())
	return cmd.Run()
}

func (g *GitAdapter) Update(dir string) error {
	pullCmd := exec.Command("git", "-C", dir, "pull")
	fetchCmd := exec.Command("git", "-C", dir, "fetch")

	if err := pullCmd.Run(); err != nil {
		return err
	}
	return fetchCmd.Run()
}
