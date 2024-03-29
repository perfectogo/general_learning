package encodingjson

type Order struct {
	ID          string `json:"id"`
	DateOrdered string `json:"date_ordered"`
	CustomerID  string `json:"customer_id"`
	Items       []Item `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
