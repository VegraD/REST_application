package combine

import (
	"Assignment-1/convert"
	"Assignment-1/structs"
)

/*
A function for combining a slice of universities with a slice of countries.
Parameters:

	unis: A slice of university structs.
	countries: A slice of country structs.

Returns a slice of UniAndCountry structs.
*/
func CombineUniAndCountry(unis []structs.University, countries []structs.Country) []structs.UniAndCountry {
	var outputs []structs.UniAndCountry

	for _, i := range unis {
		for _, j := range countries {
			if i.Country == j.Name["common"] {
				outputs = append(outputs, convert.ToUniAndCountry(i, j))
			}
		}

	}
	return outputs
}
