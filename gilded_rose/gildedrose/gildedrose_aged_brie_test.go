package gildedrose_test

import (
	"gilded_rose/gildedrose"
	"testing"
)

func Test_UpdateQuality_AgedBrie_QualityIncreases(t *testing.T) {
	items := gildedrose.CreateSingleItem("Aged Brie", 2, 0)
	gildedrose.UpdateQuality(items)
	if items[0].Quality != 1 {
		t.Errorf("Expected quality to be 1, got %d", items[0].Quality)
	}
}
