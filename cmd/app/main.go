package main

import (
	"context"
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/scraping"
	service "mercurio-web-scraping/internal/domain/services"
	"mercurio-web-scraping/internal/infra/mongodb"
	"mercurio-web-scraping/internal/infra/repositories"
)

func main() {
	config := config.GetConfig()
	ctx := context.Background()

	mongodb := mongodb.GetConnection(ctx, config)
	mongodb.SeedDB()

	linkRepo := repositories.NewLinkMongoRepository(mongodb.DB)
	harvestRepo := repositories.NewHarvestMongoRepository(mongodb.DB)

	linkService := service.NewLinkService(linkRepo)
	harvestService := service.NewHarvestService(harvestRepo)

	scraping := scraping.NewScraping(linkService, harvestService)
	scraping.Start(ctx)
}
