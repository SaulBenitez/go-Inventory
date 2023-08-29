package dtos

type RegisterProductRequest struct {
	Name		string  `json:"name" validate:"required,min=3,max=50"`
	Description string  `json:"description" validate:"required,min=3,max=100"`
	Price       float64 `json:"price" validate:"required,min=1.0"`
	UserEmail	string  `json:"user_email" validate:"required,email"`
}