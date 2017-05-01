package yelp

// Category describes a business.
type Category struct {
	Alias string `json:"alias"`
	Title string `json:"title"`
}

// Business defines a business returned by the Yelp API.
type Business struct {
	Categories   []Category  `json:"categories"`
	Coodinates   Coordinates `json:"coordinates"`
	DisplayPhone string      `json:"display_phone"`
	Distance     float64     `json:"distance"`
	ID           string      `json:"id"`
	ImageURL     string      `json:"image_url"`
	IsClosed     bool        `json:"is_closed"`
	Location     Location    `json:"location"`
	Name         string      `json:"name"`
	Phone        string      `json:"phone"`
	Price        string      `json:"price"`
	Rating       float64     `json:"rating"`
	ReviewCount  int64       `json:"review_count"`
	URL          string      `json:"url"`
	Transactions []string    `json:"transactions"`
}
