package db

type DataBaseInterface interface {
	CloseConnection()
	GetActiveSlugs(userId string) ([]string, error)
	DeleteSlug(slug string) error
	InsertSlug(slug string) error
	UpdateUserSlugs(userId string, insertSlugs, deleteSlugs []string) error
}
