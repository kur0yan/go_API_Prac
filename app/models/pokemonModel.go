package models

type Pokemon struct {
	pokedex_number int32    `json:"pokedex_id,omitempty" bson:"pokedex_number,omitempty"`
	generation     string   `json:"generation,omitempty" bson:"generation,omitempty"`
	name           string   `json:"name,omitempty" bson:"name,omitempty"`
	classification string   `json:"classification,omitempty" bson:"classification,omitempty"`
	poke_type      []string `json:"types,omitempty" bson:"type,omitempty"`
	height         float32  `json:"height_in_m,omitempty" bson:"height_m,omitempty"`
	weight         float32  `json:"weight_in_kg,omitempty" bson:"weight_kg,omitempty"`
	hp             int32    `json:"hp,omitempty" bson:"hp,omitempty"`
	attack         int32    `json:"attack,omitempty" bson:"attack,omitempty"`
	defense        int32    `json:"defense,omitempty" bson:"defense,omitempty"`
	sp_attack      int32    `json:"special_attack,omitempty" bson:"sp_attack,omitempty"`
	sp_def         int32    `json:"special_defense,omitempty" bson:"sp_defense,omitempty"`
	speed          int32    `json:"pokedex_Id,omitempty" bson:"pokedex_number,omitempty"`
	bst            int32    `json:"base_total,omitempty" bson:"base_total,omitempty"`
	isLegendary    bool     `json:"is_legendary,omitempty" bson:"is_legendary,omitempty"`
}
