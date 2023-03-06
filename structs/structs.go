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
	CCA2      string                 `json:"cca2"`
	CCA3      string                 `json:"cca3"`
}

type Diagnostic struct {
	UniAPI     string `json:"universitiesapi"`
	CountryAPI string `json:"countriesapi"`
	Version    string `json:"version"`
	Uptime     string `json:"uptime"`
}
