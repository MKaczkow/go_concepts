package main

type AbandonedPlace struct {
	Name        string  `json:"name"`
	URL         string  `json:"place_url"`
	Details     Details `json:"details"`
	Description string  `json:"description"`
	Comments    Comment `json:"comments"`
}

type Details struct {
	Added          string      `json:"added"`
	AddedBy        string      `json:"added_by"`
	AddedByLink    string      `json:"added_by_link"`
	Accessibility  string      `json:"accessibility"`
	Attractiveness string      `json:"attractiveness"`
	Gallery        string      `json:"gallery"`
	GalleryLink    string      `json:"gallery_link"`
	Category       string      `json:"category"`
	CategoryLink   string      `json:"category_link"`
	Coordinates    Coordinates `json:"coordinates"`
	Location       string      `json:"location"`
	Rating         float64     `json:"rating"`
	VoteCount      int         `json:"vote_count"`
	Status         string      `json:"status"`
	Views          int         `json:"views"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	MapLink   string  `json:"map_link"`
}

type Comment struct {
	User        string `json:"user"`
	UserLink    string `json:"user_link"`
	Timestamp   string `json:"timestamp"` // Można parsować na time.Time w przyszłości
	CommentText string `json:"comment_text"`
}
