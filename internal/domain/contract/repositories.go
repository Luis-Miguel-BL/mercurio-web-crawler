package contract

import (
	"context"
	"mercurio-web-crawler/internal/domain/entities"
)

type LinkRepository interface {
	GetByUUID(context context.Context, LinkUUID string) (link entities.Link, err error)
	FindAvailableToVisits(context context.Context) (links []entities.Link, err error)
	Update(context context.Context, link entities.Link) (err error)
}

type HarvestRepository interface {
	Create(context context.Context, harvest entities.Harvest) (err error)
}
