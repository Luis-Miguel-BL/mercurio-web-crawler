package main

import (
	"context"
	"fmt"
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/entities"
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

	linkRepo := repositories.NewLinkMemoryRepository()
	harvestRepo := repositories.NewHarvestMongoRepository(mongodb.DB)

	harvest := entities.Harvest{}
	harvest.SetDefaultValues()

	harvestRepo.Create(ctx, harvest)

	linkService := service.NewLinkService(linkRepo)
	harvestService := service.NewHarvestService(harvestRepo)

	scraping := scraping.NewScraping(linkService, harvestService)

	fmt.Println("Scraping started")
	scraping.Start(ctx)
}
