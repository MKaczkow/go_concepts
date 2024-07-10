package gildedrose

import (
	"testing"
)

func TestUpdateQuality(t *testing.T) {
	// Test case 1: Normal item before sell-by date
	items := []*Item{
		{Name: "Normal Item", SellIn: 5, Quality: 10},
	}
	UpdateQuality(items)
	if items[0].SellIn != 4 {
		t.Errorf("Expected SellIn to be 4, got %d", items[0].SellIn)
	}
	if items[0].Quality != 9 {
		t.Errorf("Expected Quality to be 9, got %d", items[0].Quality)
	}

	// Test case 2: Aged Brie
	items = []*Item{
		{Name: "Aged Brie", SellIn: 2, Quality: 10},
	}
	UpdateQuality(items)
	if items[0].SellIn != 1 {
		t.Errorf("Expected SellIn to be 1, got %d", items[0].SellIn)
	}
	if items[0].Quality != 11 {
		t.Errorf("Expected Quality to be 11, got %d", items[0].Quality)
	}

	// Test case 3: Backstage passes
	items = []*Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
	}
	UpdateQuality(items)
	if items[0].SellIn != 14 {
		t.Errorf("Expected SellIn to be 14, got %d", items[0].SellIn)
	}
	if items[0].Quality != 21 {
		t.Errorf("Expected Quality to be 21, got %d", items[0].Quality)
	}

	// Test case 4: Conjured item
	items = []*Item{
		{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6},
	}
	UpdateQuality(items)
	if items[0].SellIn != 2 {
		t.Errorf("Expected SellIn to be 2, got %d", items[0].SellIn)
	}
	if items[0].Quality != 4 {
		t.Errorf("Expected Quality to be 4, got %d", items[0].Quality)
	}

	// Test case 5: Expired item
	items = []*Item{
		{Name: "Normal Item", SellIn: -1, Quality: 10},
	}
	UpdateQuality(items)
	if items[0].SellIn != -2 {
		t.Errorf("Expected SellIn to be -2, got %d", items[0].SellIn)
	}
	if items[0].Quality != 8 {
		t.Errorf("Expected Quality to be 8, got %d", items[0].Quality)
	}
}
