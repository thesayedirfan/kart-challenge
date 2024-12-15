package domain

type Order struct {
	ID       string    `json:"id"`
	Total    float64   `json:"total"`
	Discount float64   `json:"discount"`
	Items    []Item    `json:"items"`
	Products []Product `json:"products"`
}

type Item struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type OrderRequest struct {
	Items      []Item `json:"items"`
	CouponCode string `json:"couponCode"`
}
