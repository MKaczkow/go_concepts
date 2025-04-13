package models

import (
	"encoding/json"
	"testing"
)

func TestAbandonedPlaceSerialization(t *testing.T) {
	tests := []struct {
		input    AbandonedPlace
		expected string
	}{
		{
			input: AbandonedPlace{
				Name:        "Old Factory",
				URL:         "https://example.com/old-factory",
				ScrapeDate:  "2023-10-01 15:34:55",
				Description: "An old factory from the 19th century",
				Details: Details{
					Added:          "2023-01-01",
					AddedBy:        "TestUser",
					AddedByLink:    "https://example.com/user/TestUser",
					Accessibility:  "Easy",
					Attractiveness: "High",
					Gallery:        "Photo Gallery",
					GalleryLink:    "https://example.com/gallery/123",
					Category:       "Industrial",
					CategoryLink:   "https://example.com/category/industrial",
					Coordinates: Coordinates{
						Latitude:  52.4064,
						Longitude: 16.9252,
						MapLink:   "https://maps.example.com/123",
					},
					Location:  "Poznan, Poland",
					Rating:    4.5,
					VoteCount: 120,
					Status:    "Abandoned",
					Views:     5432,
				},
				Comments: []Comment{
					{
						User:        "Commenter",
						UserLink:    "https://example.com/user/Commenter",
						Timestamp:   "21 days ago",
						CommentText: "This place is amazing!",
					},
					{
						User:        "Commenter2",
						UserLink:    "https://example.com/user/Commente2",
						Timestamp:   "37 days ago",
						CommentText: "This place is terrible!",
					},
				},
				Hazards: []Hazard{
					{
						ID:          0,
						Name:        "",
						Description: "Place is sometimes visited by security guards",
						Added:       "",
						AddedBy:     "",
					},
					{
						ID:          0,
						Name:        "",
						Description: "Place is often visited by security guards",
						Added:       "",
						AddedBy:     "",
					},
				},
			},
			expected: `{
				"name": "Old Factory",
				"place_url": "https://example.com/old-factory",
				"scrape_date": "2023-10-01 15:34:55",
				"description": "An old factory from the 19th century",
				"details": {
					"added": "2023-01-01",
					"added_by": "TestUser",
					"added_by_link": "https://example.com/user/TestUser",
					"accessibility": "Easy",
					"attractiveness": "High",
					"gallery": "Photo Gallery",
					"gallery_link": "https://example.com/gallery/123",
					"category": "Industrial",
					"category_link": "https://example.com/category/industrial",
					"coordinates": {
						"latitude": 52.4064,
						"longitude": 16.9252,
						"map_link": "https://maps.example.com/123"
					},
					"location": "Poznan, Poland",
					"rating": 4.5,
					"vote_count": 120,
					"status": "Abandoned",
					"views": 5432
				},
				"comments": [
					{
					"user": "Commenter",
					"user_link": "https://example.com/user/Commenter",
					"timestamp": "21 days ago",
					"comment_text": "This place is amazing!"
					},
					{
					"user": "Commenter2",
					"user_link": "https://example.com/user/Commente2",
					"timestamp": "37 days ago",
					"comment_text": "This place is terrible!"
					}
				],
				"hazards": [
					{
					"id": 0,
					"name": "",
					"description": "Place is sometimes visited by security guards",
					"added": "",
					"added_by": ""
					},
					{
					"id": 0,
					"name": "",
					"description": "Place is often visited by security guards",
					"added": "",
					"added_by": ""
					}
				]
			}`,
		},
		{
			input: AbandonedPlace{
				Name:        "Old Factory",
				URL:         "https://example.com/old-factory",
				ScrapeDate:  "2023-10-01 15:34:55",
				Description: "An old factory from the 19th century",
				Details: Details{
					Added:          "2023-01-01",
					AddedBy:        "TestUser",
					AddedByLink:    "https://example.com/user/TestUser",
					Accessibility:  "Easy",
					Attractiveness: "High",
					Gallery:        "Photo Gallery",
					GalleryLink:    "https://example.com/gallery/123",
					Category:       "Industrial",
					CategoryLink:   "https://example.com/category/industrial",
					Coordinates: Coordinates{
						Latitude:  52.4064,
						Longitude: 16.9252,
						MapLink:   "https://maps.example.com/123",
					},
					Location:  "Poznan, Poland",
					Rating:    4.5,
					VoteCount: 120,
					Status:    "Abandoned",
					Views:     5432,
				},
				Comments: nil,
				Hazards:  nil,
			},
			expected: `{
				"name": "Old Factory",
				"place_url": "https://example.com/old-factory",
				"scrape_date": "2023-10-01 15:34:55",
				"description": "An old factory from the 19th century",
				"details": {
					"added": "2023-01-01",
					"added_by": "TestUser",
					"added_by_link": "https://example.com/user/TestUser",
					"accessibility": "Easy",
					"attractiveness": "High",
					"gallery": "Photo Gallery",
					"gallery_link": "https://example.com/gallery/123",
					"category": "Industrial",
					"category_link": "https://example.com/category/industrial",
					"coordinates": {
						"latitude": 52.4064,
						"longitude": 16.9252,
						"map_link": "https://maps.example.com/123"
					},
					"location": "Poznan, Poland",
					"rating": 4.5,
					"vote_count": 120,
					"status": "Abandoned",
					"views": 5432
				},
				"comments": null,
				"hazards": null
			}`,
		},
		{
			input: AbandonedPlace{
				Name:        "Old Factory",
				URL:         "https://example.com/old-factory",
				ScrapeDate:  "2023-10-01 15:34:55",
				Description: "An old factory from the 19th century",
				Details: Details{
					Added:          "2 months ago",
					AddedBy:        "TestUser",
					AddedByLink:    "",
					Accessibility:  "Easy",
					Attractiveness: "High",
					Gallery:        "",
					GalleryLink:    "",
					Category:       "Industrial",
					CategoryLink:   "",
					Coordinates: Coordinates{
						Latitude:  52.4064,
						Longitude: 16.9252,
						MapLink:   "",
					},
					Location:  "Poznan, Poland",
					Rating:    4.5,
					VoteCount: 0,
					Status:    "Active",
					Views:     0,
				},
				Comments: nil,
				Hazards:  nil,
			},
			expected: `{
				"name": "Old Factory",
				"place_url": "https://example.com/old-factory",
				"scrape_date": "2023-10-01 15:34:55",
				"description": "An old factory from the 19th century",
				"details": {
					"added": "2 months ago",
					"added_by": "TestUser",
					"added_by_link": "",
					"accessibility": "Easy",
					"attractiveness": "High",
					"gallery": "",
					"gallery_link": "",
					"category": "Industrial",
					"category_link": "",
					"coordinates": {
						"latitude": 52.4064,
						"longitude": 16.9252,
						"map_link": ""
					},
					"location": "Poznan, Poland",
					"rating": 4.5,
					"vote_count": 0,
					"status": "Active",
					"views": 0
				},
				"comments": null,
				"hazards": null
			}`,
		},
	}

	for _, tt := range tests {
		// Convert the struct to JSON
		jsonData, err := json.MarshalIndent(tt.input, "", "  ")
		if err != nil {
			t.Fatalf("Failed to marshal AbandonedPlace: %v", err)
		}

		// For comparing JSON, we need to convert our expected string to a normalized form
		var expectedJSON map[string]interface{}
		var actualJSON map[string]interface{}

		// Unmarshal the expected JSON string into a map
		if err := json.Unmarshal([]byte(tt.expected), &expectedJSON); err != nil {
			t.Fatalf("Failed to unmarshal expected JSON: %v", err)
		}

		// Unmarshal our actual JSON into a map for comparison
		if err := json.Unmarshal(jsonData, &actualJSON); err != nil {
			t.Fatalf("Failed to unmarshal actual JSON: %v", err)
		}

		// Re-marshal both to get normalized JSON strings for comparison
		expectedBytes, _ := json.Marshal(expectedJSON)
		actualBytes, _ := json.Marshal(actualJSON)

		// Compare the normalized JSON strings
		if string(expectedBytes) != string(actualBytes) {
			t.Errorf("JSON output mismatch: \nGot: %s\nWant: %s", string(jsonData), tt.expected)
		}
	}
}

func TestAbandonedPlaceDeserialization(t *testing.T) {
	tests := []struct {
		input    string
		expected AbandonedPlace
	}{
		{
			input: `{
				"name": "Old Factory",
				"place_url": "https://example.com/old-factory",
				"scrape_date": "2023-10-01 15:34:55",
				"description": "An old factory from the 19th century",
				"details": {
					"added": "2023-01-01",
					"added_by": "TestUser",
					"added_by_link": "https://example.com/user/TestUser",
					"accessibility": "Easy",
					"attractiveness": "High",
					"gallery": "Photo Gallery",
					"gallery_link": "https://example.com/gallery/123",
					"category": "Industrial",
					"category_link": "https://example.com/category/industrial",
					"coordinates": {
						"latitude": 52.4064,
						"longitude": 16.9252,
						"map_link": "https://maps.example.com/123"
					},
					"location": "Poznan, Poland",
					"rating": 4.5,
					"vote_count": 120,
					"status": "Abandoned",
					"views": 5432
				},
				"comments": [
					{
					"user": "Commenter",
					"user_link": "https://example.com/user/Commenter",
					"timestamp": "21 days ago",
					"comment_text": "This place is amazing!"
					},
					{
					"user": "Commenter2",
					"user_link": "https://example.com/user/Commente2",
					"timestamp": "37 days ago",
					"comment_text": "This place is terrible!"
					}
				],
				"hazards": [
					{
					"id": 0,
					"name": "",
					"description": "Place is sometimes visited by security guards",
					"added": "",
					"added_by": ""
					},
					{
					"id": 0,
					"name": "",
					"description": "Place is often visited by security guards",
					"added": "",
					"added_by": ""
					}
				]
			}`,
			expected: AbandonedPlace{
				Name:        "Old Factory",
				URL:         "https://example.com/old-factory",
				ScrapeDate:  "2023-10-01 15:34:55",
				Description: "An old factory from the 19th century",
				Details: Details{
					Added:          "2023-01-01",
					AddedBy:        "TestUser",
					AddedByLink:    "https://example.com/user/TestUser",
					Accessibility:  "Easy",
					Attractiveness: "High",
					Gallery:        "Photo Gallery",
					GalleryLink:    "https://example.com/gallery/123",
					Category:       "Industrial",
					CategoryLink:   "https://example.com/category/industrial",
					Coordinates: Coordinates{
						Latitude:  52.4064,
						Longitude: 16.9252,
						MapLink:   "https://maps.example.com/123",
					},
					Location:  "Poznan, Poland",
					Rating:    4.5,
					VoteCount: 120,
					Status:    "Abandoned",
					Views:     5432,
				},
				Comments: []Comment{
					{
						User:        "Commenter",
						UserLink:    "https://example.com/user/Commenter",
						Timestamp:   "21 days ago",
						CommentText: "This place is amazing!",
					},
					{
						User:        "Commenter2",
						UserLink:    "https://example.com/user/Commente2",
						Timestamp:   "37 days ago",
						CommentText: "This place is terrible!",
					},
				},
				Hazards: []Hazard{
					{
						ID:          0,
						Name:        "",
						Description: "Place is sometimes visited by security guards",
						Added:       "",
						AddedBy:     "",
					},
					{
						ID:          0,
						Name:        "",
						Description: "Place is often visited by security guards",
						Added:       "",
						AddedBy:     "",
					},
				},
			},
		},
		{
			input: `{
				"name": "Old Factory",
				"place_url": "https://example.com/old-factory",
				"scrape_date": "2023-10-01 15:34:55",
				"description": "An old factory from the 19th century",
				"details": {
					"added": "2023-01-01",
					"added_by": "TestUser",
					"added_by_link": "https://example.com/user/TestUser",
					"accessibility": "Easy",
					"attractiveness": "High",
					"gallery": "Photo Gallery",
					"gallery_link": "https://example.com/gallery/123",
					"category": "Industrial",
					"category_link": "https://example.com/category/industrial",
					"coordinates": {
						"latitude": 52.4064,
						"longitude": 16.9252,
						"map_link": "https://maps.example.com/123"
					},
					"location": "Poznan, Poland",
					"rating": 4.5,
					"vote_count": 120,
					"status": "Abandoned",
					"views": 5432
				},
				"comments": null,
				"hazards": null
			}`,
			expected: AbandonedPlace{
				Name:        "Old Factory",
				URL:         "https://example.com/old-factory",
				ScrapeDate:  "2023-10-01 15:34:55",
				Description: "An old factory from the 19th century",
				Details: Details{
					Added:          "2023-01-01",
					AddedBy:        "TestUser",
					AddedByLink:    "https://example.com/user/TestUser",
					Accessibility:  "Easy",
					Attractiveness: "High",
					Gallery:        "Photo Gallery",
					GalleryLink:    "https://example.com/gallery/123",
					Category:       "Industrial",
					CategoryLink:   "https://example.com/category/industrial",
					Coordinates: Coordinates{
						Latitude:  52.4064,
						Longitude: 16.9252,
						MapLink:   "https://maps.example.com/123",
					},
					Location:  "Poznan, Poland",
					Rating:    4.5,
					VoteCount: 120,
					Status:    "Abandoned",
					Views:     5432,
				},
				Comments: nil,
				Hazards:  nil,
			},
		},
		{
			input: `{
				"name": "Old Factory",
				"place_url": "https://example.com/old-factory",
				"scrape_date": "2023-10-01 15:34:55",
				"description": "An old factory from the 19th century",
				"details": {
					"added": "2 months ago",
					"added_by": "TestUser",
					"added_by_link": "",
					"accessibility": "Easy",
					"attractiveness": "High",
					"gallery": "",
					"gallery_link": "",
					"category": "Industrial",
					"category_link": "",
					"coordinates": {
						"latitude": 52.4064,
						"longitude": 16.9252,
						"map_link": ""
					},
					"location": "Poznan, Poland",
					"rating": 4.5,
					"vote_count": 0,
					"status": "Active",
					"views": 0
				},
				"comments": null,
				"hazards": null
			}`,
			expected: AbandonedPlace{
				Name:        "Old Factory",
				URL:         "https://example.com/old-factory",
				ScrapeDate:  "2023-10-01 15:34:55",
				Description: "An old factory from the 19th century",
				Details: Details{
					Added:          "2 months ago",
					AddedBy:        "TestUser",
					AddedByLink:    "",
					Accessibility:  "Easy",
					Attractiveness: "High",
					Gallery:        "",
					GalleryLink:    "",
					Category:       "Industrial",
					CategoryLink:   "",
					Coordinates: Coordinates{
						Latitude:  52.4064,
						Longitude: 16.9252,
						MapLink:   "",
					},
					Location:  "Poznan, Poland",
					Rating:    4.5,
					VoteCount: 0,
					Status:    "Active",
					Views:     0,
				},
				Comments: nil,
				Hazards:  nil,
			},
		},
	}

	for _, tt := range tests {
		var place AbandonedPlace
		err := json.Unmarshal([]byte(tt.input), &place)

		if err != nil {
			t.Fatalf("Failed to unmarshal AbandonedPlace: %v", err)
		}

		// Check name and basic fields
		if place.Name != tt.expected.Name {
			t.Errorf("Name mismatch: got %v, want %v", place.Name, tt.expected.Name)
		}
		if place.URL != tt.expected.URL {
			t.Errorf("URL mismatch: got %v, want %v", place.URL, tt.expected.URL)
		}
		if place.ScrapeDate != tt.expected.ScrapeDate {
			t.Errorf("ScrapeDate mismatch: got %v, want %v", place.ScrapeDate, tt.expected.ScrapeDate)
		}
		if place.Description != tt.expected.Description {
			t.Errorf("Description mismatch: got %v, want %v", place.Description, tt.expected.Description)
		}

		// Check Details field
		if place.Details.Added != tt.expected.Details.Added {
			t.Errorf("Details.Added mismatch: got %v, want %v", place.Details.Added, tt.expected.Details.Added)
		}
		if place.Details.AddedBy != tt.expected.Details.AddedBy {
			t.Errorf("Details.AddedBy mismatch: got %v, want %v", place.Details.AddedBy, tt.expected.Details.AddedBy)
		}
		if place.Details.Rating != tt.expected.Details.Rating {
			t.Errorf("Details.Rating mismatch: got %v, want %v", place.Details.Rating, tt.expected.Details.Rating)
		}

		// Check Coordinates
		if place.Details.Coordinates.Latitude != tt.expected.Details.Coordinates.Latitude {
			t.Errorf("Coordinates.Latitude mismatch: got %v, want %v", place.Details.Coordinates.Latitude, tt.expected.Details.Coordinates.Latitude)
		}
		if place.Details.Coordinates.Longitude != tt.expected.Details.Coordinates.Longitude {
			t.Errorf("Coordinates.Longitude mismatch: got %v, want %v", place.Details.Coordinates.Longitude, tt.expected.Details.Coordinates.Longitude)
		}

		// Check Category
		if place.Details.Category != tt.expected.Details.Category {
			t.Errorf("Category mismatch: got %v, want %v", place.Details.Category, tt.expected.Details.Category)
		}
		if place.Details.CategoryLink != tt.expected.Details.CategoryLink {
			t.Errorf("CategoryLink mismatch: got %v, want %v", place.Details.CategoryLink, tt.expected.Details.CategoryLink)
		}
		// Check Gallery
		if place.Details.Gallery != tt.expected.Details.Gallery {
			t.Errorf("Gallery mismatch: got %v, want %v", place.Details.Gallery, tt.expected.Details.Gallery)
		}
		if place.Details.GalleryLink != tt.expected.Details.GalleryLink {
			t.Errorf("GalleryLink mismatch: got %v, want %v", place.Details.GalleryLink, tt.expected.Details.GalleryLink)
		}
		// Check Location
		if place.Details.Location != tt.expected.Details.Location {
			t.Errorf("Location mismatch: got %v, want %v", place.Details.Location, tt.expected.Details.Location)
		}
		// Check Accessibility
		if place.Details.Accessibility != tt.expected.Details.Accessibility {
			t.Errorf("Accessibility mismatch: got %v, want %v", place.Details.Accessibility, tt.expected.Details.Accessibility)
		}
		// Check Attractiveness
		if place.Details.Attractiveness != tt.expected.Details.Attractiveness {
			t.Errorf("Attractiveness mismatch: got %v, want %v", place.Details.Attractiveness, tt.expected.Details.Attractiveness)
		}
		// Check Status
		if place.Details.Status != tt.expected.Details.Status {
			t.Errorf("Status mismatch: got %v, want %v", place.Details.Status, tt.expected.Details.Status)
		}
		// Check Views
		if place.Details.Views != tt.expected.Details.Views {
			t.Errorf("Views mismatch: got %v, want %v", place.Details.Views, tt.expected.Details.Views)
		}
		// Check VoteCount
		if place.Details.VoteCount != tt.expected.Details.VoteCount {
			t.Errorf("VoteCount mismatch: got %v, want %v", place.Details.VoteCount, tt.expected.Details.VoteCount)
		}
		// Check MapLink
		if place.Details.Coordinates.MapLink != tt.expected.Details.Coordinates.MapLink {
			t.Errorf("MapLink mismatch: got %v, want %v", place.Details.Coordinates.MapLink, tt.expected.Details.Coordinates.MapLink)
		}
		// Check Comments
		if len(place.Comments) != len(tt.expected.Comments) {
			t.Errorf("Comments length mismatch: got %d, want %d", len(place.Comments), len(tt.expected.Comments))
		} else {
			for i, comment := range place.Comments {
				if comment.User != tt.expected.Comments[i].User {
					t.Errorf("Comment[%d].User mismatch: got %v, want %v", i, comment.User, tt.expected.Comments[i].User)
				}
				if comment.CommentText != tt.expected.Comments[i].CommentText {
					t.Errorf("Comment[%d].CommentText mismatch: got %v, want %v", i, comment.CommentText, tt.expected.Comments[i].CommentText)
				}
			}
		}
		// Check Hazards
		if len(place.Hazards) != len(tt.expected.Hazards) {
			t.Errorf("Hazards length mismatch: got %d, want %d", len(place.Hazards), len(tt.expected.Hazards))
		} else {
			for i, hazard := range place.Hazards {
				if hazard.Description != tt.expected.Hazards[i].Description {
					t.Errorf("Hazard[%d].Description mismatch: got %v, want %v", i, hazard.Description, tt.expected.Hazards[i].Description)
				}
			}
		}
	}
}
