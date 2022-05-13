package contract

import (
	"context"
	"mercurio-web-scraping/internal/domain/entities"
)

type LinkService interface {
	GetByUUID(context context.Context, LinkUUID string) (link entities.Link, err error)
	FindAvailableToVisits(context context.Context) (links []entities.Link, err error)
	Update(context context.Context, link entities.Link) (err error)
}

type HarvestService interface {
	Create(context context.Context, harvest entities.Harvest) (err error)
}
