package encodingjson

import (
	"encoding/json"
	"fmt"
	"time"
)

type RFC822ZTime struct {
	time.Time
}

func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := rt.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}

func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
	if err != nil {
		return err
	}
	*rt = RFC822ZTime{t}
	return nil
}

func CustomParsing() {
	data := `
	{
		"id": "12345",
		"items": [
			{
				"id": "xyz123",
				"name": "Thing 1"
			},
			{
				"id": "abc789",
				"name": "Thing 2"
			}
		],
		"date_ordered": "01 May 20 13:01 +0000",
		"customer_id": "3"
	}`

	var o Order
	err := json.Unmarshal([]byte(data), &o)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%+v\n", o)
	fmt.Println(o.DateOrdered)
	
	out, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
