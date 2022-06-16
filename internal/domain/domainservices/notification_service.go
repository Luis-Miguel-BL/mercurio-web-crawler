package domainservices

import (
	"context"
	"mercurio-web-scraping/internal/domain/contract"
	"mercurio-web-scraping/internal/domain/entities"
)

type NotificationService struct {
	repo contract.NotificationRepository
}

func NewNotificationService(repo contract.NotificationRepository) *NotificationService {
	return &NotificationService{
		repo: repo,
	}
}

func (s NotificationService) FindByTarget(ctx context.Context, target entities.HarvestType) (notifications []entities.Notification, err error) {
	return s.repo.FindByTarget(ctx, target)
}
