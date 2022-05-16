package scraping

import (
	"context"
	"log"
	"mercurio-web-scraping/internal/domain/link_handlers"
	service "mercurio-web-scraping/internal/domain/services"
	"time"
)

type Scraping struct {
	ctx      context.Context
	svc      service.Service
	handlers link_handlers.LinkHandlers
}

func NewScraping(ctx context.Context, svc service.Service, handlers link_handlers.LinkHandlers) *Scraping {
	return &Scraping{ctx: ctx, svc: svc, handlers: handlers}
}

func (s *Scraping) Start(ctx context.Context) {
	log.Println("Starting Scraping...")
scrapingLoop:
	for {
		select {
		case <-ctx.Done():
			log.Println("Breaking Scraping...")
			break scrapingLoop
		case <-time.After(2 * time.Second):
			linksToVisit, err := s.svc.LinkService.FindAvailableToVisits(ctx)
			if err != nil {
				panic("cannot be find available links to visit")
			}

			if len(linksToVisit) > 0 {
				for _, linkToVisit := range linksToVisit {
					handleLink, ok := s.handlers[linkToVisit.Slug]
					if !ok {
						log.Println("link handle not found ", linkToVisit.Slug)
						continue
					}
					log.Println("scraping link: ", linkToVisit.Slug)
					handleLink.HandlerLink(linkToVisit)
				}
			}
		}
	}

}
