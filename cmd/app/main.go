package main

import (
	"context"
	"fmt"
	"mercurio-web-crawler/internal/domain/crawler"
	service "mercurio-web-crawler/internal/domain/services"
	"mercurio-web-crawler/internal/infra/repositories"
)

func main() {
	ctx := context.Background()

	linkRepo := repositories.NewLinkMemoryRepository()
	harvestRepo := repositories.NewHarvestMemoryRepository()

	linkService := service.NewLinkService(linkRepo)
	harvestService := service.NewHarvestService(harvestRepo)

	crawler := crawler.NewCrawler(linkService, harvestService)

	fmt.Println("Crawler started")
	crawler.Start(ctx)
}
