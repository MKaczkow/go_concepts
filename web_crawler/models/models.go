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

// UserProfile represents the overall structure of a user's profile page.
type User struct {
	Username         string `json:"username"`
	JoinedDate       string `json:"joined_date"` // Use time.Time for actual parsing
	LastLogin        string `json:"last_login"`  // Use time.Time for actual parsing
	LastLoginTooltip string `json:"last_login_tooltip"`
	AvatarURL        string `json:"avatar_url"`
	// TODO: maybe later to create cross-reference with places
	// VisitedPlaces     []PlaceSummary `json:"visited_places"`
	// AddedPlaces       []PlaceSummary `json:"added_places"`
	// Comments     []Comment `json:"comments"`
	CommentsCount int    `json:"comments_count"`
	ChangesCount  int    `json:"changes_count"` // Count of changes made
	ScrapeDate    string `json:"scrape_date"`   // Date when the profile was scraped
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

// Define the hazard types as constants or a map for easy lookup
var HazardTypes = map[string]Hazard{
	"Wystepujace promieniowanie":        {ID: 0, Name: "Wystepujace promieniowanie"},
	"Niebezpieczne materiały":           {ID: 1, Name: "Niebezpieczne materiały"},
	"Aktywność paranormalna":            {ID: 2, Name: "Aktywność paranormalna"},
	"Miejsce chronione":                 {ID: 3, Name: "Miejsce chronione"},
	"Miejsce monitorowane":              {ID: 4, Name: "Miejsce monitorowane"},
	"Wysokie ryzyko upadku z wysokości": {ID: 5, Name: "Wysokie ryzyko upadku z wysokości"},
	"Wysokie ryzyko zawalenia":          {ID: 6, Name: "Wysokie ryzyko zawalenia"},
	"Nieznane":                          {ID: 7, Name: "Nieznane"},
}
