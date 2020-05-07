package entity

type Publisher struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	URL     string `json:"url"`
}

type allPublishers []Publisher

var Publishers = allPublishers{

	{
		ID:      1,
		Name:    "Tony Stark",
		Address: "Sea Breeze",
		Phone:   "555-3332",
		URL:     "tonystark.io",
	},

	{
		ID:      54,
		Name:    "Daisy Johnson",
		Address: "Jamrock Ave",
		Phone:   "554-2453",
		URL:     "daisycakes.net",
	},
}
