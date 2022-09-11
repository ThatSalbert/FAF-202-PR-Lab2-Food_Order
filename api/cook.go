package api

type cook struct {
	rank         int    `json:"rank"`
	proficiency  int    `json:"proficiency"`
	name         string `json:"name"`
	catch_phrase string `json:"catch-phrase"`
}

var cook1 = cook{rank: 3, proficiency: 4, name: "Cook1", catch_phrase: "Cook1"}
var cook2 = cook{rank: 2, proficiency: 3, name: "Cook2", catch_phrase: "Cook2"}
var cook3 = cook{rank: 2, proficiency: 2, name: "Cook3", catch_phrase: "Cook3"}
var cook4 = cook{rank: 1, proficiency: 2, name: "Cook4", catch_phrase: "Cook4"}

var cooks = []cook{cook1, cook2, cook3, cook4}
