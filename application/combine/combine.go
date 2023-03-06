package combine

import (
	"Assignment-1/application/convert"
	"Assignment-1/structs"
)

/*
A function for combining a slice of universities with a slice of countries.
Parameters:

	unis: A slice of university structs.
	countries: A slice of country structs.
	limit: how many unis per country to report

Returns a slice of UniAndCountry structs.
*/
func CombineUniAndCountry(unis []structs.University, countries []structs.Country, limit int) []structs.UniAndCountry {
	var outputs []structs.UniAndCountry
	for _, i := range unis {
		for _, j := range countries {
			if i.Country == j.Name["common"] {
				if limit != 0 {
					if len(outputs) >= limit {
						break
					}
				}
				outputs = append(outputs, convert.ToUniAndCountry(i, j))
			}
		}
		if limit != 0 {
			if len(outputs) >= limit {
				break
			}
		}
	}
	return outputs
}
