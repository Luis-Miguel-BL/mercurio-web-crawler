package service

import (
	"context"
	"mercurio-web-crawler/internal/domain/contract"
	"mercurio-web-crawler/internal/domain/entities"
	"time"
)

type LinkService struct {
	repo contract.LinkRepository
}

func NewLinkService(repo contract.LinkRepository) *LinkService {
	return &LinkService{
		repo: repo,
	}
}

func (s *LinkService) GetByUUID(context context.Context, LinkUUID string) (link entities.Link, err error) {
	return s.repo.GetByUUID(context, LinkUUID)
}
func (s *LinkService) FindAvailableToVisits(context context.Context) (links []entities.Link, err error) {
	return s.repo.FindAvailableToVisits(context)
}
func (s *LinkService) Update(context context.Context, link entities.Link) (err error) {
	_, err = s.repo.GetByUUID(context, link.Base.UUID)
	if err != nil {
		return err
	}

	link.UpdatedAt = time.Now()
	return s.repo.Update(context, link)
}
