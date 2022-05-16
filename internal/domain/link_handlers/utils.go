package link_handlers

import "strings"

func buildAddress(
	ZipCode string,
	City string,
	Street string,
	StreetNumber string,
	Neighborhood string,
	PoisList []string,
	Complement string,
) string {
	return Street + " nº " + StreetNumber + " Bairro:" + Neighborhood + " " + City + " " + Complement + strings.Join(PoisList, ", ")

}
