package item

type Cook struct {
	Rank         int    `json:"rank"`
	Proficiency  int    `json:"proficiency"`
	Name         string `json:"name"`
	Catch_phrase string `json:"catch-phrase"`
}

var cook1 = Cook{Rank: 3, Proficiency: 4, Name: "Cook1", Catch_phrase: "Cook1"}
var cook2 = Cook{Rank: 2, Proficiency: 3, Name: "Cook2", Catch_phrase: "Cook2"}
var cook3 = Cook{Rank: 2, Proficiency: 2, Name: "Cook3", Catch_phrase: "Cook3"}
var cook4 = Cook{Rank: 1, Proficiency: 2, Name: "Cook4", Catch_phrase: "Cook4"}

var Cooks = []Cook{cook1, cook2, cook3, cook4}
