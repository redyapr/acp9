package product

type ProductResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data []Product `json:"data"`
}
