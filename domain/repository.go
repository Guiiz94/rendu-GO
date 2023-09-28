package domain

type Repository struct {
	Name       string
	HTMLURL    string
	UpdatedAt  string
	CloneURL   string
}

type RepositoryLister interface {
	List(username string) ([]Repository, error)
}

type CSVWriter interface {
	Write(repos []Repository) error
}

type GitManager interface {
	Clone(repo Repository, dir string) error
	Update(dir string) error
}

type Zipper interface {
	Zip(dir, output string) error
}
