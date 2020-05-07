package entity

type ShoppingBasket struct {
	ID       int      `json:"id"`
	Customer Customer `json:"customer"`
	BookISBN string   `json:"book_isbn"`
	Count    int      `json:"count"`
}

type shoppingBasket []ShoppingBasket

var Basket = shoppingBasket{

	{
		ID: 4,
		Customer: Customer{
			ID:      6,
			Email:   "fitz@hotmail.com",
			Name:    "Leopold",
			Phone:   "353-4294",
			Address: "Grove",
		},
		BookISBN: "758492",
		Count:    2,
	},

	{
		ID: 67,
		Customer: Customer{
			ID:      2,
			Email:   "yoyo@hotmail.com",
			Name:    "Elena",
			Phone:   "343-0594",
			Address: "Robinson Road",
		},
		BookISBN: "438594",
		Count:    1,
	},
}
