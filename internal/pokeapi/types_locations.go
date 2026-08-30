package pokeapi

type Shallow struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Locations struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Shallow `json:"results"`
}

type Location struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod Shallow `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int     `json:"rate"`
			Version Shallow `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location Shallow `json:"location"`
	Names    []struct {
		Name     string  `json:"name"`
		Language Shallow `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        Shallow `json:"pokemon"`
		VersionDetails []struct {
			Version          Shallow `json:"version"`
			MaxChance        int     `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int     `json:"min_level"`
				MaxLevel        int     `json:"max_level"`
				ConditionValues []any   `json:"condition_values"`
				Chance          int     `json:"chance"`
				Method          Shallow `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
