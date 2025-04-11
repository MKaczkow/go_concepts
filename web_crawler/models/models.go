package models

type AbandonedPlace struct {
	Name        string    `json:"name"`
	URL         string    `json:"place_url"`
	ScrapeDate  string    `json:"scrape_date"`
	Description string    `json:"description"`
	Details     Details   `json:"details"`
	Comments    []Comment `json:"comments"`
	Hazards     []Hazard  `json:"hazards"`
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

type Category struct {
	ID   int
	Name string
}

// TODO: var or const?
var (
	Railway             = Category{0, "Kolejowe"}
	Other               = Category{1, "Inne"}
	HousesMansionsFlats = Category{2, "Domy, dworki i bloki"}
	Industrial          = Category{3, "Industrialne"}
	Military            = Category{4, "Militarne"}
	CastlesMonuments    = Category{5, "Zamki i Zabytki"}
	HospitalsMedical    = Category{6, "Szpitale i obiekty medyczne"}
	UndergroundTunnels  = Category{7, "Podziemia i tunele"}
	ResortsHotels       = Category{8, "Ośrodki wypoczynkowe i hotele"}
	ShoppingCenters     = Category{9, "Centra handlowe i sklepy"}
)

// Almamer is good example of place having almost every danger category
// Chernobyl obviously has radiation and any HTML contains .js function to choose hazard marker
type Hazard struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Added       string `json:"added"`
	AddedBy     string `json:"added_by"`
}

// TODO: var or const?
// var (
// 	Radiation          = Hazard{0, "Wystepujace promieniowanie"}
// 	DangerousMaterials = Hazard{1, "Niebezpieczne materiały"}
// 	ParanormalActivity = Hazard{2, "Aktywność paranormalna"}
// 	SiteGuarded        = Hazard{3, "Miejsce chronione"}
// 	SiteMonitored      = Hazard{4, "Miejsce monitorowane"}
// 	RiskOfFalling      = Hazard{5, "Wysokie ryzyko upadku z wysokości"}
// 	RiskOfCollapse     = Hazard{6, "Wysokie ryzyko zawalenia"}
// )
