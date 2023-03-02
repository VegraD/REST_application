package DB

import (
	"Assignment-1/structs"
)

var Db = make(map[int]structs.UniAndCountry)

func InitStudentsStorage() map[int]structs.UniAndCountry {

	// Prepopulate with data
	uni1 := structs.UniAndCountry{
		Name:     "NTNU",
		Country:  "Norway",
		Isocode:  "NO",
		Webpages: []string{"http://www.ntnu.no/", "http://www.google.no/"},
		Languages: map[string]string{"nno": "Norwegian Nynorsk",
			"nob": "Norwegian Bokmål",
			"smi": "Sami"},
		Map: "https://www.openstreetmap.org/relation/2978650",
	}

	uni2 := structs.UniAndCountry{
		Name:      "University College of Arts, Crafts and Design",
		Country:   "Sweden",
		Isocode:   "SE",
		Webpages:  []string{"http://www.hig.se/"},
		Languages: map[string]string{"swe": "Swedish"},
		Map:       "https://www.openstreetmap.org/relation/52822",
	}

	uni3 := structs.UniAndCountry{
		Name:     "University of Vaasa",
		Country:  "Finland",
		Isocode:  "FI",
		Webpages: []string{"http://www.uwasa.fi/"},
		Languages: map[string]string{"fin": "Finnish",
			"swe": "Swedish"},
		Map: "openstreetmap.org/relation/54224",
	}

	Db[0] = uni1
	Db[1] = uni2
	Db[2] = uni3

	return Db
}
