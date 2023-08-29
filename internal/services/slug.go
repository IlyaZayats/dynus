package services

import (
	"fmt"
	"github.com/IlyaZayats/dynus/internal/entity"
	"github.com/IlyaZayats/dynus/internal/interfaces"
	"strconv"
)

type SlugService struct {
	repo interfaces.SlugRepository
}

func NewSlugService(repo interfaces.SlugRepository) (*SlugService, error) {
	return &SlugService{
		repo: repo,
	}, nil
}

func (s *SlugService) InsertSlug(slug, slugChance string) error {
	chance, _ := strconv.ParseFloat(slugChance, 64)
	return s.repo.InsertSlug(entity.Slug{
		Name:   slug,
		Chance: chance,
	})
}

func (s *SlugService) DeleteSlug(slug string) error {
	return s.repo.DeleteSlug(entity.Slug{Name: slug})
}

func (s *SlugService) GetActiveSlugs(userId string) ([]string, error) {
	id, _ := strconv.Atoi(userId)
	return s.repo.GetActiveSlugs(entity.User{Id: id})
}

func (s *SlugService) UpdateUserSlugs(userId string, insertSlugs, deleteSlugs []string, ttl map[string]string) error {
	id, _ := strconv.Atoi(userId)
	return s.repo.UpdateUserSlugs(entity.User{Id: id}, insertSlugs, deleteSlugs, ttl)
}

func (s *SlugService) GetHistory(date string) ([]string, error) {
	history, err := s.repo.GetHistory(date)
	var csvArray []string
	var str string
	if err != nil {
		return csvArray, err
	}
	for i, v := range history {
		str = fmt.Sprintf("%s;%s;%s;%s\n", v.UserId, v.Slug, v.Operation, v.Timestamp)
		if i == len(history)-1 {
			str = fmt.Sprintf("%s;%s;%s;%s", v.UserId, v.Slug, v.Operation, v.Timestamp)
		}
		csvArray = append(csvArray, str)
	}
	return csvArray, nil

}
