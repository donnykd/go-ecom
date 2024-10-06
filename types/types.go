package types

import "time"

type UserStorage interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RegisterUser struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=100"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ProductStorage interface {
	GetProducts() ([]Product, error)
	CreateProduct(CreateProduct) error
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateProduct struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity" validate:"required"`
}
