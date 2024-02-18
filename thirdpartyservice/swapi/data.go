package swapi

type People struct {
	Name      string   `json:"name"`
	Height    string   `json:"height,omitempty"`
	Mass      string   `json:"mass,omitempty"`
	HairColor string   `json:"hair_color,omitempty"`
	SkinColor string   `json:"skinColor,omitempty"`
	EyeColor  string   `json:"skin_color,omitempty"`
	BirthYear string   `json:"birth_year,omitempty"`
	Gender    string   `json:"gender,omitempty"`
	Films     []string `json:"films,omitempty"`
	Vehicles  []string `json:"vehicles,omitempty"`
	Created   string   `json:"created,omitempty"`
	Edited    string   `json:"edited,omitempty"`
	URL       string   `json:"url"`
}

type PeopleData struct {
	Count      int      `json:"count"`
	Next       string   `json:"next"`
	Previous   string   `json:"previous"`
	ResultList []People `json:"results"`
}

type Film struct {
	Title       string  `json:"title"`
	EpisodeID   int     `json:"episode_id"`
	Director    *string `json:"director,omitempty"`
	Producer    *string `json:"producer,omitempty"`
	ReleaseDate *string `json:"release_date,omitempty"`
}

type Vehicle struct {
	Name  string `json:"name"`
	Model string `json:"model"`
}
