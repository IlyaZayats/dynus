package interfaces

import "github.com/IlyaZayats/dynus/internal/entity"

type SlugRepository interface {
	GetActiveSlugs(user entity.User) ([]string, error)
	DeleteSlug(slug entity.Slug) error
	InsertSlug(slug entity.Slug) error
	UpdateUserSlugs(user entity.User, insertSlugs, deleteSlugs []string, ttl map[string]string) error

	CleanupByTTL() error
	GetHistory(date string) ([]entity.History, error)
}
