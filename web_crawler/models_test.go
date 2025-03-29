package main

import (
	"encoding/json"
	"testing"
)

func TestAbandonedPlaceDeserialization(t *testing.T) {
	jsonData := `{
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
		"description": "An old factory from the 19th century",
		"comments": {
			"user": "Commenter",
			"user_link": "https://example.com/user/Commenter",
			"timestamp": "2023-02-15 14:30",
			"comment_text": "This place is amazing!"
		}
	}`

	var place AbandonedPlace
	err := json.Unmarshal([]byte(jsonData), &place)

	if err != nil {
		t.Fatalf("Failed to unmarshal AbandonedPlace: %v", err)
	}

	// Test AbandonedPlace fields
	if place.Description != "An old factory from the 19th century" {
		t.Errorf("Expected description '%s', got '%s'", "An old factory from the 19th century", place.Description)
	}

	// Test Details fields
	details := place.Details
	if details.Added != "2023-01-01" {
		t.Errorf("Expected Added '%s', got '%s'", "2023-01-01", details.Added)
	}
	if details.AddedBy != "TestUser" {
		t.Errorf("Expected AddedBy '%s', got '%s'", "TestUser", details.AddedBy)
	}
	if details.Rating != 4.5 {
		t.Errorf("Expected Rating %f, got %f", 4.5, details.Rating)
	}
	if details.VoteCount != 120 {
		t.Errorf("Expected VoteCount %d, got %d", 120, details.VoteCount)
	}

	// Test Coordinates fields
	coords := details.Coordinates
	if coords.Latitude != 52.4064 {
		t.Errorf("Expected Latitude %f, got %f", 52.4064, coords.Latitude)
	}
	if coords.Longitude != 16.9252 {
		t.Errorf("Expected Longitude %f, got %f", 16.9252, coords.Longitude)
	}

	// Test Comment fields
	comment := place.Comments
	if comment.User != "Commenter" {
		t.Errorf("Expected User '%s', got '%s'", "Commenter", comment.User)
	}
	if comment.CommentText != "This place is amazing!" {
		t.Errorf("Expected CommentText '%s', got '%s'", "This place is amazing!", comment.CommentText)
	}
}

func TestAbandonedPlaceSerialization(t *testing.T) {
	place := AbandonedPlace{
		Details: Details{
			Added:          "2023-03-15",
			AddedBy:        "Author",
			AddedByLink:    "https://example.com/user/Author",
			Accessibility:  "Medium",
			Attractiveness: "Medium",
			Gallery:        "Photos",
			GalleryLink:    "https://example.com/gallery/456",
			Category:       "Military",
			CategoryLink:   "https://example.com/category/military",
			Coordinates: Coordinates{
				Latitude:  54.3520,
				Longitude: 18.6466,
				MapLink:   "https://maps.example.com/456",
			},
			Location:  "Gdansk, Poland",
			Rating:    3.8,
			VoteCount: 45,
			Status:    "Partially Demolished",
			Views:     2345,
		},
		Description: "An abandoned military bunker",
		Comments: Comment{
			User:        "Visitor",
			UserLink:    "https://example.com/user/Visitor",
			Timestamp:   "2023-04-01 10:15",
			CommentText: "Difficult to access but worth it",
		},
	}

	jsonData, err := json.Marshal(place)
	if err != nil {
		t.Fatalf("Failed to marshal AbandonedPlace: %v", err)
	}

	// Now unmarshal it back to verify
	var newPlace AbandonedPlace
	err = json.Unmarshal(jsonData, &newPlace)
	if err != nil {
		t.Fatalf("Failed to unmarshal serialized AbandonedPlace: %v", err)
	}

	// Verify the unmarshaled data matches original
	if newPlace.Description != place.Description {
		t.Errorf("Description mismatch after serialization")
	}
	if newPlace.Details.Rating != place.Details.Rating {
		t.Errorf("Rating mismatch after serialization")
	}
	if newPlace.Details.Coordinates.Latitude != place.Details.Coordinates.Latitude {
		t.Errorf("Latitude mismatch after serialization")
	}
	if newPlace.Comments.CommentText != place.Comments.CommentText {
		t.Errorf("Comment text mismatch after serialization")
	}
}
