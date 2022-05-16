package link_handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mercurio-web-scraping/internal/domain/entities"
	service "mercurio-web-scraping/internal/domain/services"
	"net/http"
	"time"
)

type ZapImoveisHandler struct {
	ctx context.Context
	svc service.Service
}

type zapImoveisHarvestInfo struct {
	Name           string   `json:"name"`
	Link           string   `json:"link"`
	CountSuites    int64    `json:"count_suites"`
	CountBathrooms int64    `json:"count_bathrooms"`
	CountBedrooms  int64    `json:"count_bedrooms"`
	Pricing        string   `json:"pricing"`
	Address        string   `json:"address"`
	Contact        string   `json:"contact"`
	Medias         []string `json:"medias"`
}

type ZapImoveisResponse struct {
	Search struct {
		Result struct {
			Listings []struct {
				Listing struct {
					UpdatedAt    time.Time `json:"updated_at"`
					Suites       []int64   `json:"suites"`
					Bathrooms    []int64   `json:"bathrooms"`
					Bedrooms     []int64   `json:"bedrooms"`
					PricingInfos []struct {
						YearlyIptu      string `json:"yearly_iptu"`
						Price           string `json:"price"`
						BusinessType    string `json:"business_type"`
						MonthlyCondoFee string `json:"monthly_condo_fee"`
					} `json:"pricingInfos"`
					OriginalAddress struct {
						ZipCode      string   `json:"zip_code"`
						City         string   `json:"city"`
						Neighborhood string   `json:"neighborhood"`
						PoisList     []string `json:"pois_list"`
						Complement   string   `json:"complement"`
						Street       string   `json:"street"`
						StreetNumber string   `json:"streetNumber"`
					} `json:"originalAddress"`
					AdvertiserContact struct {
						Phones []string `json:"phones"`
					} `json:"advertiserContact"`
					WhatsappNumber string `json:"whatsappNumber"`
				} `json:"listing"`
				Medias []struct {
					Id  string `json:"id"`
					Url string `json:"url"`
				} `json:"medias"`
				Link struct {
					Name string `json:"name"`
					Href string `json:"href"`
				} `json:"link"`
			} `json:"listings"`
		} `json:"result"`
	} `json:"search"`
}

func BuildZapImoveisHandler(ctx context.Context, svc service.Service) *ZapImoveisHandler {
	return &ZapImoveisHandler{ctx: ctx, svc: svc}
}

func (h *ZapImoveisHandler) HandlerLink(link entities.Link) {
	req, err := http.NewRequest(http.MethodGet, link.Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-domain", "www.zapimoveis.com.br")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var bodyJson ZapImoveisResponse
	json.Unmarshal(bodyBytes, &bodyJson)

	for _, result := range bodyJson.Search.Result.Listings {
		medias := make([]string, 0)
		for _, media := range result.Medias {
			url := fmt.Sprintf("https://resizedimgs.zapimoveis.com.br/crop/420x236/vr.images.sp/%s.jpg", media.Id)
			medias = append(medias, url)
		}
		resultURL := "https://www.zapimoveis.com.br" + result.Link.Href

		zapImoveisHarvestInfo := zapImoveisHarvestInfo{
			Link:           resultURL,
			Name:           result.Link.Name,
			CountSuites:    result.Listing.Suites[0],
			CountBathrooms: result.Listing.Bathrooms[0],
			CountBedrooms:  result.Listing.Bedrooms[0],
			Pricing:        "R$ " + result.Listing.PricingInfos[0].Price,
			Address:        buildAddress(result.Listing.OriginalAddress.ZipCode, result.Listing.OriginalAddress.City, result.Listing.OriginalAddress.Street, result.Listing.OriginalAddress.StreetNumber, result.Listing.OriginalAddress.Neighborhood, result.Listing.OriginalAddress.PoisList, result.Listing.OriginalAddress.Complement),
			Contact:        "Telefone: " + result.Listing.AdvertiserContact.Phones[0] + " - WhatsApp: " + result.Listing.WhatsappNumber,
			Medias:         medias,
		}
		rawDataBytes, rawDataErr := json.Marshal(result)
		harvestInfoBytes, harvestInfoErr := json.Marshal(zapImoveisHarvestInfo)

		harvestExist, findHarvestErr := h.svc.HarvestService.FindByPageLink(h.ctx, resultURL)
		if err != nil {
			log.Printf("Error  %+v", err)
			continue
		}

		if rawDataErr != nil || harvestInfoErr != nil || findHarvestErr != nil {
			log.Printf("Error  %+v %+v %+v", rawDataErr, harvestInfoErr, findHarvestErr)
			link.SetErrorVisit()
		}
		rawData := string(rawDataBytes)
		harvestInfo := string(harvestInfoBytes)
		if harvestExist.UUID == "" {
			linkHarvest := link.CreateHarvest(rawData, resultURL, harvestInfo)
			h.svc.HarvestService.Create(h.ctx, linkHarvest)
		} else if harvestExist.RawData != rawData {
			harvestExist.RawData = rawData
			harvestExist.Info = harvestInfo
			err := h.svc.HarvestService.Update(h.ctx, harvestExist)
			if err != nil && err.Error() != "mongo: no documents in result" {
				log.Printf("Error  %+v", err)
				link.SetErrorVisit()
			}
		}

	}
	link.SetVisit()
	h.svc.LinkService.Update(h.ctx, link)

}
