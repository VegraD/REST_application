package convert

import (
	"Assignment-1/structs"
)

/*
Function that takes a university struct and a country struct, combining the two into a UniAndCountry struct.
Parameters:

	uni: University struct to be combined
	country: Country struct to be combined

Returns a UniAndCountry struct created by combining the parameter structs.
*/
func ToUniAndCountry(uni structs.University, country structs.Country) structs.UniAndCountry {
	var newUni structs.UniAndCountry

	newUni.Name = uni.Name
	newUni.Country = uni.Country
	newUni.Isocode = uni.AlphaTwoCode
	newUni.Webpages = uni.WebPages
	newUni.Languages = country.Languages
	newUni.Map = country.Maps["openStreetMaps"]

	return newUni
}
