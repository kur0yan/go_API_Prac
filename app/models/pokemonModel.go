package models

type Pokemon struct {
	Pokedex_number int32    `json:"pokedex_id,omitempty" bson:"pokedex_number,omitempty"`
	Generation     string   `json:"generation,omitempty" bson:"generation,omitempty"`
	Name           string   `json:"name,omitempty" bson:"name,omitempty"`
	Classification string   `json:"classification,omitempty" bson:"classification,omitempty"`
	Poke_type      []string `json:"types,omitempty" bson:"type,omitempty"`
	Hp             int32    `json:"hp,omitempty" bson:"hp,omitempty"`
	Attack         int32    `json:"attack,omitempty" bson:"attack,omitempty"`
	Defense        int32    `json:"defense,omitempty" bson:"defense,omitempty"`
	Sp_attack      int32    `json:"special_attack,omitempty" bson:"sp_attack,omitempty"`
	Sp_def         int32    `json:"special_defense,omitempty" bson:"sp_defense,omitempty"`
	Speed          int32    `json:"speed,omitempty" bson:"speed,omitempty"`
	Bst            int32    `json:"base_total,omitempty" bson:"base_total,omitempty"`
	IsLegendary    bool     `json:"is_legendary,omitempty" bson:"is_legendary,omitempty"`
}
