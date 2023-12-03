package encodingjson

import (
	"encoding/json"
	"fmt"
)

func Ex() error {
	var o Order

	data := `
		{
			"id" : "0f9abf7a-b41e-46ed-8a39-18de1ef04ebe", 
			"date_ordered" : "2006-01-02 15:04:05",
			"customer_id" : "c7ceaa8b-0bba-4822-a534-4be5be2e2f29",
			"items" : [
				{
					"id"   : "1",
					"name" : "go"
				},{
					"id"   : "2",
					"name" : "java"
				}
			]
		}
	`

	err := json.Unmarshal([]byte(data), &o)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(o)

	b, err := json.Marshal(o)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}
