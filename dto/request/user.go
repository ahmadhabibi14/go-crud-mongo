package request

type AddUserIn struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   uint8  `json:"age" validate:"required"`
}

type UpdateUserIn struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   uint8  `json:"age" validate:"required"`
}
