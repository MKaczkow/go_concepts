package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()

	feed.Add(Item{"hello", "How ya doin'?"})

	if len(feed.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()

	feed.Add(Item{"hello", "How ya doin'?"})
	feed.Add(Item{"hello", "Great, thanks!"})
	
	results := feed.GetAll()
	if len(results) != 2 {
		t.Errorf("Item was not added")
	}

}