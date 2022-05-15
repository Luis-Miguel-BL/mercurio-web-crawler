package link_handlers

import (
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/entities"
	service "mercurio-web-scraping/internal/domain/services"
)

type LinkHandler interface {
	HandlerLink(link entities.Link)
}

func GetSeedLinks(config config.Config) []entities.Link {
	return []entities.Link{
		{Url: config.ZapImoveisURL, Slug: config.ZapImoveisSlug, Origin: "ZapImoveis", Description: "Novas Casas ZapImoveis", TimeoutInSeconds: 60, HarvestType: entities.HarvestBuilding, Active: true},
	}
}

type LinkSlug = string
type LinkHandlers map[LinkSlug]LinkHandler

func GetLinkHandlers(svc service.Service) LinkHandlers {
	return map[LinkSlug]LinkHandler{config.ZapImoveisSlug: BuildZapImoveisHandler(svc)}
}
