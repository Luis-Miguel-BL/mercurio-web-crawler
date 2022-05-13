package repositories

import (
	"context"
	"errors"
	"mercurio-web-scraping/internal/domain/entities"
	"time"
)

type LinkMemoryRepository struct {
	Links []entities.Link
}

func NewLinkMemoryRepository() (repo *LinkMemoryRepository) {
	links := []entities.Link{
		{Url: "https://google.com", Slug: "teste-google", Origin: "google", Description: "teste google", TimeoutInSeconds: 5, HarvestType: entities.HarvestBuilding, Active: true},
		{Url: "https://facebook.com", Slug: "teste-facebook", Origin: "facebook", Description: "teste facebook", TimeoutInSeconds: 5, HarvestType: entities.HarvestBuilding, Active: true},
	}

	return &LinkMemoryRepository{
		Links: links,
	}
}

func (repo *LinkMemoryRepository) GetByUUID(context context.Context, LinkUUID string) (link entities.Link, err error) {
	for _, link := range repo.Links {
		if link.UUID == LinkUUID {
			return link, nil
		}
	}
	return link, errors.New("link not found")

}

func (repo *LinkMemoryRepository) FindAvailableToVisits(context context.Context) (links []entities.Link, err error) {
	for _, link := range repo.Links {
		if link.Active && link.LastVisit.Unix()+link.TimeoutInSeconds < time.Now().Unix() {
			links = append(links, link)
		}
	}
	return links, nil
}

func (repo *LinkMemoryRepository) Update(context context.Context, link entities.Link) (err error) {
	for key, repoLink := range repo.Links {
		if repoLink.UUID == link.UUID {
			repo.Links[key] = link
			return nil
		}
	}
	return errors.New("link not found")
}
