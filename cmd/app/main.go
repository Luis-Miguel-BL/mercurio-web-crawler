package main

import (
	"context"
	"fmt"
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/crawler"
	service "mercurio-web-scraping/internal/domain/services"
	"mercurio-web-scraping/internal/infra/mongodb"
	"mercurio-web-scraping/internal/infra/repositories"
)

func main() {
	config := config.GetConfig()
	ctx := context.Background()

	mongodb := mongodb.Connection(ctx, config)

	linkRepo := repositories.NewLinkMemoryRepository()
	harvestRepo := repositories.NewHarvestMongoRepository(mongodb)

	linkService := service.NewLinkService(linkRepo)
	harvestService := service.NewHarvestService(harvestRepo)

	crawler := crawler.NewCrawler(linkService, harvestService)

	fmt.Println("Crawler started")
	crawler.Start(ctx)
}
