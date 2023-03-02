package structs

// University and country to display to user.
type UniAndCountry struct {
	Name      string            `json:"name,omitempty"`
	Country   string            `json:"country,omitempty"`
	Isocode   string            `json:"isocode,omitempty"`
	Webpages  []string          `json:"webpages,omitempty"`
	Languages map[string]string `json:"languages,omitempty"`
	Map       string            `json:"map,omitempty"`
}

// University to create when fetching from third-party
type University struct {
	Name         string   `json:"name"`
	Country      string   `json:"country"`
	AlphaTwoCode string   `json:"alpha_two_code"`
	WebPages     []string `json:"web_pages"`
}

// Country to create when fetching from third-party
type Country struct {
	Name      map[string]interface{} `json:"name"`
	Languages map[string]string      `json:"languages"`
	Borders   []string               `json:"borders"`
	Maps      map[string]string      `json:"maps"`
}

func ToUniAndCountry(country Country, uni University) UniAndCountry {
	var newUni UniAndCountry

	//TODO: Either implement marshalling or combining country and university info

	return newUni

}
