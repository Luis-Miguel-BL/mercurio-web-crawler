package main

import (
	"context"
	"mercurio-web-scraping/internal/application/appservices"
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/domainservices"
	"mercurio-web-scraping/internal/domain/link_handlers"
	"mercurio-web-scraping/internal/domain/scraping"
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
	notificationRepo := repositories.NewNotificationMongoRepository(mongodb.DB)

	app_svc := appservices.GetService(config)
	domain_svc := domainservices.GetServices(*app_svc, linkRepo, harvestRepo, notificationRepo)

	handlers := link_handlers.GetLinkHandlers(ctx, *domain_svc, *app_svc)
	scraping := scraping.NewScraping(ctx, *domain_svc, handlers)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		scraping.Start(ctx)
	}()
	wg.Wait()
}
