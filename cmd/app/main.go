package main

import (
	"context"
	"fmt"
	"mercurio-web-crawler/internal/config"
	"mercurio-web-crawler/internal/domain/crawler"
	service "mercurio-web-crawler/internal/domain/services"
	"mercurio-web-crawler/internal/infra/mongodb"
	"mercurio-web-crawler/internal/infra/repositories"
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
