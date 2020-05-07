package entity

type Customer struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type allCustomers []Customer

var Customers = allCustomers{

	{
		ID:      2,
		Email:   "yoyo@hotmail.com",
		Name:    "Elena",
		Phone:   "343-0594",
		Address: "Robinson Road",
	},

	{
		ID:      6,
		Email:   "fitz@hotmail.com",
		Name:    "Leopold",
		Phone:   "353-4294",
		Address: "Grove",
	},
}
