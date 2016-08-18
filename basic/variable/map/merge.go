package main

import (
	"fmt"
)

func main() {
	map_1 := map[string]interface{}{
		"string": "string",
		"int":    1,
		"map": map[string]interface{}{
			"string_1": "sub_string",
		},
	}

	map_2 := map[string]interface{}{
		"string":  "string_2",
		"int":     2,
		"float64": 3.14159,
		"map": map[string]interface{}{
			"string_2": "sub_string",
		},
	}

	map_3 := map[string]interface{}{
		"string":  "string_3",
		"int":     3,
		"float64": 3.14,
		"map": map[string]interface{}{
			"string_3": "sub_string",
		},
	}
	map_all := merge_map(map_1, map_2)
	fmt.Println(map_all)
	map_all = merge_map(map_1, map_2, map_3)
	fmt.Println(map_all)
}

func merge_map(maps ...map[string]interface{}) map[string]interface{} {
	return nil
}
