package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"

	"github.com/Guiiz94/rendu-GO/adapter"
	"github.com/Guiiz94/rendu-GO/useCase"
	// "github.com/Guiiz94/rendu-GO/domain"
)

const baseDir = "./repositories"
const outputPath = "./repositories/repositories.zip"
const limit = 100

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	username := os.Getenv("GIT_PSEUDO")
	if username == "" {
		log.Fatal("Please set the USERNAME environment variable")
	}

	githubAdapter := &adapters.GithubAPIAdapter{}
	csvAdapter := &adapters.CSVAdapter{}
	gitAdapter := &adapters.GitAdapter{}
	zipAdapter := &adapters.ZipAdapter{}

	listUC := &usecases.ListRepositoriesUseCase{RepoLister: githubAdapter}
	cloneUC := &usecases.CloneRepositoriesUseCase{Git: gitAdapter}
	saveUC := &usecases.SaveToCSVUseCase{Writer: csvAdapter}
	updateUC := &usecases.UpdateRepositoriesUseCase{Git: gitAdapter}
	zipUC := &usecases.ZipRepositoriesUseCase{Zipper: zipAdapter}

	// username := "Guiiz94" 

	repos, err := listUC.List("Guiiz94", limit)
	if err != nil {
		log.Fatalf("Error listing repositories: %v", err)
	}

	err = saveUC.Save(repos)
	if err != nil {
		log.Fatalf("Error saving repositories to CSV: %v", err)
	}

	err = cloneUC.CloneAll(repos, baseDir)
	if err != nil {
		log.Fatalf("Error cloning repositories: %v", err)
	}

	err = updateUC.UpdateAll(repos, baseDir)
	if err != nil {
		log.Fatalf("Error updating repositories: %v", err)
	}

	err = zipUC.ZipAll(baseDir, outputPath)
	if err != nil {
		log.Fatalf("Error zipping repositories: %v", err)
	}

	fmt.Println("Operation completed successfully!")

	httpServer := &adapters.HTTPServerAdapter{ZipFilePath: outputPath}
	httpServer.Start()

}
