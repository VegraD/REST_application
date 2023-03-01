package main

import (
	"Assignment-1/structs"
	"log"
)

func InitStudentsStorage() map[string]structs.UniAndCountry {
	db := make(map[string]structs.UniAndCountry)

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
	s2 := Student{Name: "Berit", Age: 20, StudentID: "2"}
	s3 := Student{Name: "Ole", Age: 21, StudentID: "3"}
	db. += uni1
	db.Add(s1)
	db.Add(s2)
	db.Add(s3)
	log.Println("Prepopulated DB ...")

	return &db
}
