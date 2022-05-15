package config

type Config struct {
	MongoURI       string
	MongoDBName    string
	ZapImoveisURL  string
	ZapImoveisSlug string
}

var MongoURI = "mongodb://localhost:27017"
var MongoDBName = "mercurio-web-scraping"
var ZapImoveisURL = "https://glue-api.zapimoveis.com.br/v2/listings?categoryPage=RESULT&business=SALE&listingType=USED&unitTypesV3=HOME&unitSubTypes=UnitSubType_NONE,TWO_STORY_HOUSE,SINGLE_STOREY_HOUSE,KITNET&unitTypes=HOME&usageTypes=RESIDENTIAL&text=Casa&sort=updatedAt+ASC&addressState=Minas+Gerais&addressCity=Par%C3%A1+de+Minas"
var ZapImoveisSlug = "house-zapimoveis"

func GetConfig() Config {
	return Config{MongoURI: MongoURI, MongoDBName: MongoDBName, ZapImoveisURL: ZapImoveisURL, ZapImoveisSlug: ZapImoveisSlug}
}
