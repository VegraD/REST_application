package convert

import (
	"Assignment-1/structs"
)

func ToUniAndCountry(uni structs.University, country structs.Country) structs.UniAndCountry {
	var newUni structs.UniAndCountry

	newUni.Name = uni.Name
	newUni.Country = uni.Country
	newUni.Isocode = uni.AlphaTwoCode
	newUni.Webpages = uni.WebPages
	newUni.Languages = country.Languages
	newUni.Map = country.Maps["openStreetMaps"]
	//TODO: Error handling

	return newUni
}
