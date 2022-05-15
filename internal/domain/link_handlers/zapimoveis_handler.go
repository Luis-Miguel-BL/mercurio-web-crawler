package link_handlers

import (
	"fmt"
	"mercurio-web-scraping/internal/domain/entities"
	service "mercurio-web-scraping/internal/domain/services"
	"time"
)

type ZapImoveisHandler struct {
	svc service.Service
}

func BuildZapImoveisHandler(svc service.Service) *ZapImoveisHandler {
	return &ZapImoveisHandler{svc: svc}
}

func (h *ZapImoveisHandler) HandlerLink(link entities.Link) {

	fmt.Printf("handle link %+v", link)
	time.Sleep(time.Second * 5)
	fmt.Printf("handle link end -----%+v", link)
}
