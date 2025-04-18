package gildedrose_test

import (
	"gilded_rose/gildedrose"
	"testing"
)

func Test_UpdateQuality_Sulfuras_QualityDoesntChange(t *testing.T) {
	items := gildedrose.CreateSingleItem("Sulfuras, Hand of Ragnaros", 0, 80)
	gildedrose.UpdateQuality(items)
	if items[0].Quality != 80 {
		t.Errorf("Expected quality to be 80, got %d", items[0].Quality)
	}
}
