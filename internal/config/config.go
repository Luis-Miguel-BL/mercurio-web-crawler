package config

type Config struct {
	MongoURI       string
	MongoDBName    string
	ZapImoveisURL  string
	ZapImoveisSlug string
}

var MongoURI = "mongodb://localhost:27017"
var MongoDBName = "mercurio-web-scraping"
var ZapImoveisURL = "https://glue-api.zapimoveis.com.br/v2/listings?categoryPage=RESULT&sort=updatedAt+ASC&addressState=Minas+Gerais&addressCity=Par%C3%A1+de+Minas&size=350"
var ZapImoveisSlug = "building-zapimoveis"

func GetConfig() Config {
	return Config{MongoURI: MongoURI, MongoDBName: MongoDBName, ZapImoveisURL: ZapImoveisURL, ZapImoveisSlug: ZapImoveisSlug}
}
