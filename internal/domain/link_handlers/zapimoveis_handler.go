package link_handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mercurio-web-scraping/internal/application/appservices"
	"mercurio-web-scraping/internal/domain/domainservices"
	"mercurio-web-scraping/internal/domain/entities"
	"net/http"
	"time"
)

type ZapImoveisHandler struct {
	ctx       context.Context
	domainSVC domainservices.Service
	appSVC    appservices.Service
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
type ZapImoveisLiting struct {
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
}
type ZapImoveisResponse struct {
	Search struct {
		Result struct {
			Listings []ZapImoveisLiting `json:"listings"`
		} `json:"result"`
	} `json:"search"`
}

func BuildZapImoveisHandler(ctx context.Context, domainSVC domainservices.Service, appSVC appservices.Service) *ZapImoveisHandler {
	return &ZapImoveisHandler{ctx: ctx, domainSVC: domainSVC, appSVC: appSVC}
}

func (h *ZapImoveisHandler) HandlerLink(link entities.Link) {
	result, err := makeRequest(link.Url)
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range result.Search.Result.Listings {
		tmpHarvest, err := buildHarvest(result)
		if err != nil {
			log.Printf("Error buildHarvest  %+v", err)
			link.SetErrorVisit()
		}
		harvestExist, err := h.domainSVC.HarvestService.FindByPageLink(h.ctx, tmpHarvest.PageLink)
		if err != nil && error.Error(err) != "mongo: no documents in result" {
			log.Printf("Error FindByPageLink  %+v", err)
			link.SetErrorVisit()
		}

		if harvestExist.UUID == "" {
			linkHarvest := link.CreateHarvest(tmpHarvest.RawData, tmpHarvest.PageLink, tmpHarvest.Info, entities.HarvestBuilding)
			h.domainSVC.HarvestService.Create(h.ctx, linkHarvest)
			h.domainSVC.NotificationService.FindAndNotifyByTargets(h.ctx, linkHarvest)

		} else if harvestExist.RawData != tmpHarvest.RawData {
			harvestExist.RawData = tmpHarvest.RawData
			harvestExist.Info = tmpHarvest.Info
			err := h.domainSVC.HarvestService.Update(h.ctx, harvestExist)
			if err != nil {
				log.Printf("Error  %+v", err)
				link.SetErrorVisit()
			}
		}

	}
	link.SetVisit()
	h.domainSVC.LinkService.Update(h.ctx, link)
}

func makeRequest(url string) (response ZapImoveisResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return response, err
	}
	req.Header.Set("x-domain", "www.zapimoveis.com.br")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var bodyJson ZapImoveisResponse
	json.Unmarshal(bodyBytes, &bodyJson)

	return bodyJson, err
}

func buildHarvest(data ZapImoveisLiting) (harvest entities.Harvest, err error) {
	medias := make([]string, 0)
	for _, media := range data.Medias {
		url := fmt.Sprintf("https://resizedimgs.zapimoveis.com.br/crop/420x236/vr.images.sp/%s.jpg", media.Id)
		medias = append(medias, url)
	}
	resultURL := "https://www.zapimoveis.com.br" + data.Link.Href

	var CountSuites int64
	var CountBathrooms int64
	var CountBedrooms int64
	if len(data.Listing.Suites) > 0 {
		CountSuites = data.Listing.Suites[0]
	}
	if len(data.Listing.Bathrooms) > 0 {
		CountBathrooms = data.Listing.Bathrooms[0]
	}
	if len(data.Listing.Bedrooms) > 0 {
		CountBedrooms = data.Listing.Bedrooms[0]
	}

	priceDescription := map[string]string{
		"SALE":   "",
		"RENTAL": "Por mês",
	}

	zapImoveisHarvestInfo := zapImoveisHarvestInfo{
		Link:           resultURL,
		Name:           data.Link.Name,
		CountSuites:    CountSuites,
		CountBathrooms: CountBathrooms,
		CountBedrooms:  CountBedrooms,
		Pricing:        "R$ " + data.Listing.PricingInfos[0].Price + " " + priceDescription[data.Listing.PricingInfos[0].BusinessType],
		Address:        buildAddress(data.Listing.OriginalAddress.ZipCode, data.Listing.OriginalAddress.City, data.Listing.OriginalAddress.Street, data.Listing.OriginalAddress.StreetNumber, data.Listing.OriginalAddress.Neighborhood, data.Listing.OriginalAddress.PoisList, data.Listing.OriginalAddress.Complement),
		Contact:        "Telefone: " + data.Listing.AdvertiserContact.Phones[0] + " - WhatsApp: " + data.Listing.WhatsappNumber,
		Medias:         medias,
	}
	rawDataBytes, err := json.Marshal(data)
	if err != nil {
		return harvest, err
	}
	harvestInfoBytes, err := json.Marshal(zapImoveisHarvestInfo)
	if err != nil {
		return harvest, err
	}

	rawData := string(rawDataBytes)
	harvestInfo := string(harvestInfoBytes)

	harvest.RawData = rawData
	harvest.Info = harvestInfo
	harvest.PageLink = resultURL

	return harvest, err
}
