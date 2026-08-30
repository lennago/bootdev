package pokeapi

type Pokemon struct {
	Pokemon pokemonAPI
	Species PokemonSpecies
}

type pokemonAPI struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_defaul"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool    `json:"is_hidden"`
		Slot     int     `json:"slot"`
		Ability  Shallow `json:"ability"`
	} `json:"abilities"`
	Forms       []Shallow `json:"forms"`
	GameIndices []struct {
		GameIndex int     `json:"game_index"`
		Version   Shallow `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item           Shallow `json:"item"`
		VersionDetails []struct {
			Rarity  int     `json:"rarity"`
			Version Shallow `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move                Shallow `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int     `json:"level_learned_at"`
			VersionGroup    Shallow `json:"version_group"`
			MoveLearnMethod Shallow `json:"move_learn_method"`
			Order           int     `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	SpeciesSh Shallow `json:"species"`
	Sprites   struct {
		BackDefault      string `json:"back_default"`
		BackFemale       any    `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  any    `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           string `json:"back_default"`
					BackShiny             string `json:"back_shiny"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					BackTransparent       string `json:"back_transparent"`
					FrontDefault          string `json:"front_default"`
					FrontShiny            string `json:"front_shiny"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
					FrontTransparent      string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default"`
						BackFemale       any    `json:"back_female"`
						BackShiny        string `json:"back_shiny"`
						BackShinyFemale  any    `json:"back_shiny_female"`
						FrontDefault     string `json:"front_default"`
						FrontFemale      any    `json:"front_female"`
						FrontShiny       string `json:"front_shiny"`
						FrontShinyFemale any    `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"lates"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int     `json:"base_stat"`
		Effort   int     `json:"effort"`
		Stat     Shallow `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int     `json:"slot"`
		Type Shallow `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation Shallow `json:"generation"`
		Types      []struct {
			Slot int     `json:"slot"`
			Type Shallow `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
	PastAbilities []struct {
		Generation Shallow `json:"generation"`
		Abilities  []struct {
			IsHidden bool    `json:"is_hidden"`
			Slot     int     `json:"slot"`
			Ability  Shallow `json:"ability"`
		} `json:"abilities"`
	} `json:"past_abilities"`
}

type PokemonSpecies struct {
	ID                   int     `json:"id"`
	Name                 string  `json:"name"`
	Order                int     `json:"order"`
	GenderRate           int     `json:"gender_rate"`
	CaptureRate          int     `json:"capture_rate"`
	BaseHappiness        int     `json:"base_happiness"`
	IsBaby               bool    `json:"is_baby"`
	IsLegendary          bool    `json:"is_legendary"`
	IsMythical           bool    `json:"is_mythical"`
	HatchCounter         int     `json:"hatch_counter"`
	HasGenderDifferences bool    `json:"has_gender_differences"`
	FormsSwitchable      bool    `json:"forms_switchable"`
	GrowthRate           Shallow `json:"growth_rate"`
	PokedexNumbers       []struct {
		EntryNumber int     `json:"entry_number"`
		Pokedex     Shallow `json:"pokedex"`
	} `json:"pokedex_numbers"`
	EggGroups          []Shallow `json:"egg_groups"`
	Color              Shallow   `json:"color"`
	Shape              Shallow   `json:"shape"`
	EvolvesFromSpecies Shallow   `json:"evolves_from_species"`
	EvolutionChain     struct {
		URL string `json:"url"`
	} `json:"evolution_chain"`
	Habitat    Shallow `json:"habitat"`
	Generation Shallow `json:"generation"`
	Names      []struct {
		Name     string  `json:"name"`
		Language Shallow `json:"language"`
	} `json:"names"`
	FlavorTextEntries []struct {
		FlavorText string  `json:"flavor_text"`
		Language   Shallow `json:"language"`
		Version    Shallow `json:"version"`
	} `json:"flavor_text_entries"`
	FormDescriptions []struct {
		Description string  `json:"description"`
		Language    Shallow `json:"language"`
	} `json:"form_descriptions"`
	Genera []struct {
		Genus    string  `json:"genus"`
		Language Shallow `json:"language"`
	} `json:"genera"`
	Varieties []struct {
		IsDefault bool    `json:"is_default"`
		Pokemon   Shallow `json:"pokemon"`
	} `json:"varieties"`
}
