package scraping

import (
	"context"
	"fmt"
	"mercurio-web-scraping/internal/domain/contract"
	"time"
)

type Scraping struct {
	linkService    contract.LinkService
	harvestService contract.HarvestService
}

func NewScraping(linkService contract.LinkService, harvestService contract.HarvestService) *Scraping {
	return &Scraping{linkService: linkService, harvestService: harvestService}
}

func (c *Scraping) Start(context context.Context) {
	for {
		linksToVisit, err := c.linkService.FindAvailableToVisits(context)
		if err != nil {
			panic("cannot be find available links to visit")
		}

		if len(linksToVisit) > 0 {
			for _, linkToVisit := range linksToVisit {
				fmt.Printf("\n visitou o link %+v\n ", linkToVisit)

			}
		}

		time.Sleep(time.Second * 5)
	}

}
