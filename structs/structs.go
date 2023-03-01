package structs

type UniAndCountry struct {
	Name      string            `json:"name,omitempty"`
	Country   string            `json:"country,omitempty"`
	Isocode   string            `json:"isocode,omitempty"`
	Webpages  []string          `json:"webpages,omitempty"`
	Languages map[string]string `json:"languages,omitempty"` //TODO: make map[string]string for correct output
	Map       string            `json:"map,omitempty"`
}
