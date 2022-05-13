package crawler

import (
	"context"
	"fmt"
	"mercurio-web-scraping/internal/domain/contract"
	"time"
)

type Crawler struct {
	linkService    contract.LinkService
	harvestService contract.HarvestService
}

func NewCrawler(linkService contract.LinkService, harvestService contract.HarvestService) *Crawler {
	return &Crawler{linkService: linkService, harvestService: harvestService}
}

func (c *Crawler) Start(context context.Context) {
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
