package user

type UserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type UserResponseSingle struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UserResponseLogin struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    UserLogin `json:"data"`
}
