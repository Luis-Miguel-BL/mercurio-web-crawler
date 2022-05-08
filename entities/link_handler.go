package entities

type LinkHandler interface {
	HandlerLink(link Link)
}

type LinkSlug = string

var LinkHandlers map[LinkSlug]LinkHandler = map[LinkSlug]LinkHandler{}
