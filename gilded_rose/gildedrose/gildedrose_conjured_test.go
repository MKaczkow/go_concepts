package gildedrose_test

import (
	"gilded_rose/gildedrose"
	"testing"
)

func Test_UpdateQuality_Conjured_QualityDecreases(t *testing.T) {
	items := []*gildedrose.Item{
		{Name: "Conjured", SellIn: 10, Quality: 20},
	}
	expectedQuality := 18
	gildedrose.UpdateQuality(items)
	if items[0].Quality != expectedQuality {
		t.Errorf("Expected quality to be %d, got %d", expectedQuality, items[0].Quality)
	}
}

func Test_UpdateQuality_Conjured_SellInDecreases(t *testing.T) {
	items := []*gildedrose.Item{
		{Name: "Conjured", SellIn: 10, Quality: 20},
	}
	expectedSellIn := 9
	gildedrose.UpdateQuality(items)
	if items[0].SellIn != expectedSellIn {
		t.Errorf("Expected SellIn to be %d, got %d", expectedSellIn, items[0].SellIn)
	}
}

func Test_UpdateQuality_Conjured_QualityDecreasesTwiceAfterSellIn(t *testing.T) {
	items := gildedrose.CreateSingleItem("Conjured", 1, 10)
	expectedQualities := []int{8, 4, 0}
	gildedrose.UpdateQuality(items) // Day 1 (SellIn becomes 0, Quality becomes 8)
	if items[0].Quality != expectedQualities[0] {
		t.Errorf("Expected quality to be %d, got %d", expectedQualities[0], items[0].Quality)
	}
	gildedrose.UpdateQuality(items) // Day 2 (SellIn becomes -1, Quality becomes 4)
	if items[0].Quality != expectedQualities[1] {
		t.Errorf("Expected quality to be %d, got %d", expectedQualities[1], items[0].Quality)
	}
	gildedrose.UpdateQuality(items) // Day 3 (SellIn becomes -2, Quality becomes 0)
	if items[0].Quality != expectedQualities[2] {
		t.Errorf("Expected quality to be %d, got %d", expectedQualities[2], items[0].Quality)
	}
}
