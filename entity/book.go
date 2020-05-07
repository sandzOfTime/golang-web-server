package entity

type Book struct {
	ID            int    `json:"id"`
	ISBN          string `json:"isbn"`
	PublisherName string `json:"publisher_name"`
	AuthorName    string `json:"author_name"`
	AuthorAddress string `json:"author_address"`
	Title         string `json:"title"`
	Year          string `json:"year"`
	Price         string `json:"price"`
}

type allBooks []Book

var Books = allBooks{

	{
		ID:            2,
		ISBN:          "438594",
		PublisherName: "Daisy Johnson",
		AuthorName:    "Phil Coulson",
		AuthorAddress: "Mangrove Road",
		Title:         "Hunger Games",
		Year:          "1995",
		Price:         "$23.00",
	},

	{
		ID:            5,
		ISBN:          "758492",
		PublisherName: "Tony Stark",
		AuthorName:    "Maria Hill",
		AuthorAddress: "Treasure Cove",
		Title:         "The Last Dance",
		Year:          "2001",
		Price:         "$25.00",
	},
}
