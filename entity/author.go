package entity

type Author struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	URL     string `json:"url"`
}

type allAuthors []Author

var Authors = allAuthors{

	{
		ID:      4,
		Name:    "Phil Coulson",
		Address: "Mangrove Road",
		URL:     "philcoul@aol.com",
	},

	{
		ID:      67,
		Name:    "Maria Hill",
		Address: "Treasure Cove",
		URL:     "mariahill@hotmail.com",
	},
}
