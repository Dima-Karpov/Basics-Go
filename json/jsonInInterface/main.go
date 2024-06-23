package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var ks = []byte(`{
	"firstName": "Jean",
	"lastName": "Bartik",
	"age": 86,
	"education": [
		{
			"institution": "Northwest Missouri State Teachers College",
			"degree": "Bachelor of Science in Mathematics"
		},
		{
			"institution": "University of Pennsylvania",
			"degree": "Master in English"
		}
	],
	"spouse": "William Bartik",
	"children": [
		"Timothy John Bartik",
		"Jane Helen Bartik",
		"Mary Ruth Bartik"
	]
}`)

func main() {
	var f interface{}
	err := json.Unmarshal(ks, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := f.(map[string]interface{})
	fmt.Println(m["firstName"], m["lastName"])

	fmt.Println(f)
}

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string:", vv)
	case float64:
		fmt.Println("is float64:", vv)
	case []interface{}:
		fmt.Println("is an array:")
		for i, u := range vv {
			fmt.Println("index:", i, "value:", u)
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Println("index:", i, "value:", u)
			printJSON(u)
		}
	default:
		fmt.Println("I don't know about type %T", v)
	}
}
