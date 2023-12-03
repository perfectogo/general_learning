package encodingjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func EncodeDecode2() {
	const data = `
		{"name": "Fred", "age": 40}
		{"name": "Mary", "age": 21}
		{"name": "Pat", "age": 30}
	`
	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	decoder := json.NewDecoder(strings.NewReader(data))
	var b bytes.Buffer

	enc := json.NewEncoder(&b)

	for decoder.More() {
		if err := decoder.Decode(&t); err != nil {
			panic(err)
		}

		fmt.Println(t)

		if err := enc.Encode(t); err != nil {
			panic(err)
		}
	}

	out := b.String()

	fmt.Println(out)
}
