package account_role_repo

type Repo interface {
}

type Repository struct {
}

func NewRepository() Repo {
	return &Repository{}
}
