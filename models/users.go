package models

//Users Data Struct
type Users struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Phone    string `db:"phone: json:"phone:`
	Birthday string `db:"birthday" json:"birthday"`
	Status   string `db:"status" json:"status"`
}

//OrderInput data struct
type UserInput struct {
	Email    string `json:"email" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone: validate:"required:`
	Birthday string `json:"birthday" validate:"required"`
	Status   string `json:"status" validate:"required"`
}
