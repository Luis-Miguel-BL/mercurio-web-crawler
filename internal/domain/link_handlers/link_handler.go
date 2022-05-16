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

type LinkSlug = string
type LinkHandlers map[LinkSlug]LinkHandler

func GetLinkHandlers(ctx context.Context, svc service.Service) LinkHandlers {
	return map[LinkSlug]LinkHandler{config.ZapImoveisSlug: BuildZapImoveisHandler(ctx, svc)}
}
