package internal

// Nac stands for "Nacional address card", most known as DNI in argentina.
type Customer struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Nac       string `json:"nac" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Cellphone int64  `json:"cellphone" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Cvu       string `json:"cvu" validate:"required"`
	Alias     string `json:"alias" validate:"required"`
}

type Credentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
