package handler

type University struct {
	Name      string   `json:"name,omitempty"`
	Country   string   `json:"country,omitempty"`
	Isocode   string   `json:"isocode,omitempty"`
	Webpages  []string `json:"webpages,omitempty"`
	Languages []string `json:"languages,omitempty"`
	Map       string   `json:"map,omitempty"`
}
type Languages struct {

	//TODO: make language struct, include language code + full name of language

}
