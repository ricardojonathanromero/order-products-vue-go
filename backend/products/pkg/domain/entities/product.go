package entities

type Product struct {
	Id          string   `json:"id"`
	Sku         string   `json:"sku"`
	Price       float64  `json:"price"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Quantity    int64    `json:"quantity"`
	Image       string   `json:"image"`
	Categories  []string `json:"categories"`
}
