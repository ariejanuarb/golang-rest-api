package web

type ProductResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	//CreatedAt  time.Time `json:"created_at"`
	//UpdatedAt  time.Time `json:"updated_at"`
}
