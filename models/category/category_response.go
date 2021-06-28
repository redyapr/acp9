package category

type CategoryResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data []Category `json:"data"`
}

type CategoryResponseSingle struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data Category `json:"data"`
}
