package models

// Nac stands for "Nacional address card", most known as DNI in argentina.
type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Nac       string `json:"nac"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Cellphone int64  `json:"cellphone"`
	Password  string `json:"password"`
	Cvu       string `json:"cvu"`
	Alias     string `json:"alias"`
}
