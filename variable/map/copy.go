package main

import (
	"fmt"
)

func main() {
	src := map[string]interface{}{
		"string": "string",
		"int":    1,
		"map": map[string]interface{}{
			"string": "sub_string",
		},
	}
	dst := copy_map(src)

	fmt.Println(dst)
	dst["string"] = "string_1"
	test_map := dst["map"].(map[string]interface{})
	test_map["string"] = "sub_string_1"
	fmt.Println(dst)
	fmt.Println(src)
}

func copy_map(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{})
	for k, v := range src {
		switch v.(type) {
		case string:
			dst[k] = v
		case int:
			dst[k] = v
		case int64:
			dst[k] = v
		case float64:
			dst[k] = v
		default:
			src_v := v.(map[string]interface{})
			dst[k] = copy_map(src_v)
		}
	}
	return dst
}
