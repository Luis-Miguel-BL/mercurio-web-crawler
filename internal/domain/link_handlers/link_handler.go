package link_handlers

import (
	"context"
	"mercurio-web-scraping/internal/application/appservices"
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/domainservices"
	"mercurio-web-scraping/internal/domain/entities"
)

type LinkHandler interface {
	HandlerLink(link entities.Link)
}

type LinkSlug = string
type LinkHandlers map[LinkSlug]LinkHandler

func GetLinkHandlers(ctx context.Context, domainSVC domainservices.Service, appSVC appservices.Service) LinkHandlers {
	return map[LinkSlug]LinkHandler{config.ZapImoveisSlug: BuildZapImoveisHandler(ctx, domainSVC, appSVC)}
}
