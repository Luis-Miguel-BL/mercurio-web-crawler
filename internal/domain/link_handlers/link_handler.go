package link_handlers

import (
	"context"
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/entities"
	service "mercurio-web-scraping/internal/domain/services"
)

type LinkHandler interface {
	HandlerLink(link entities.Link)
}

func GetSeedLinks(config config.Config) []entities.Link {
	return []entities.Link{
		{Url: config.ZapImoveisURL, Slug: config.ZapImoveisSlug, Origin: "ZapImoveis", Description: "Novas Casas ZapImoveis", TimeoutInSeconds: 60 * 60 * 24, HarvestType: entities.HarvestBuilding, Active: true},
	}
}

type LinkSlug = string
type LinkHandlers map[LinkSlug]LinkHandler

func GetLinkHandlers(ctx context.Context, svc service.Service) LinkHandlers {
	return map[LinkSlug]LinkHandler{config.ZapImoveisSlug: BuildZapImoveisHandler(ctx, svc)}
}
