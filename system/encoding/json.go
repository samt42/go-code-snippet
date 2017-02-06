package main

import (
	"fmt"
	"encoding/json"
)
type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Sex string `json:"sex"`
	Interest interest `json:"interest"`
	Family family `json:"family"`
}

type interest struct {
	Ball []string `json:ball`
	Sports []string `json:sports`
}
type family struct {
	Son map[string]string `json:"son"`
	Daughter map[string]string `json:"daughter"`
}
func main() {
	p := person {
		Name: "Tom",
		Age: 20,
		Sex: "man",
		Interest: interest {
			Ball: []string{"basketball", "football"},
			Sports: []string{"run", "longrun"},
		},
		Family: family {
			Son: map[string]string{"son1": "Nick", "son2": "Mike"},
			Daughter: map[string]string{"dau1": "Mary", "dau2": "Lucy"},
		},
	}
	fmt.Println(p.Family.Son["son1"])
	pjson, _ := json.Marshal(p)
	fmt.Println(string(pjson))
}
