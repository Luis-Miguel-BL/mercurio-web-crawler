package main

import (
	"context"
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/link_handlers"
	"mercurio-web-scraping/internal/domain/scraping"
	service "mercurio-web-scraping/internal/domain/services"
	"mercurio-web-scraping/internal/infra/mongodb"
	"mercurio-web-scraping/internal/infra/repositories"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	config := config.GetConfig()

	mongodb := mongodb.GetConnection(ctx, config)
	mongodb.SeedDB()

	linkRepo := repositories.NewLinkMongoRepository(mongodb.DB)
	harvestRepo := repositories.NewHarvestMongoRepository(mongodb.DB)

	svc := service.GetServices(linkRepo, harvestRepo)

	handlers := link_handlers.GetLinkHandlers(ctx, *svc)
	scraping := scraping.NewScraping(ctx, *svc, handlers)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		scraping.Start(ctx)
	}()
	wg.Wait()
}
