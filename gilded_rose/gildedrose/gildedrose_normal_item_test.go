package gildedrose_test

import (
	"gilded_rose/gildedrose"
	"testing"
)

func Test_UpdateQuality_NormalItem_QualityDecreases(t *testing.T) {
	items := []*gildedrose.Item{
		{Name: "Normal", SellIn: 10, Quality: 20},
	}
	gildedrose.UpdateQuality(items)
	if items[0].Quality != 19 {
		t.Errorf("Expected quality to be 19, got %d", items[0].Quality)
	}
}

func Test_UpdateQuality_NormalItem_SellInDecreases(t *testing.T) {
	items := []*gildedrose.Item{
		{Name: "Normal", SellIn: 10, Quality: 20},
	}
	gildedrose.UpdateQuality(items)
	if items[0].SellIn != 9 {
		t.Errorf("Expected SellIn to be 9, got %d", items[0].SellIn)
	}
}

func Test_UpdateQuality_NormalItem_QualityDecreasesTwiceAfterSellIn(t *testing.T) {
	items := gildedrose.CreateSingleItem("Normal", 1, 10)
	expectedQuality := 5
	gildedrose.UpdateQuality(items) // Day 1 (SellIn becomes 0, Quality becomes 9)
	gildedrose.UpdateQuality(items) // Day 2 (SellIn becomes -1, Quality becomes 7)
	gildedrose.UpdateQuality(items) // Day 3 (SellIn becomes -2, Quality becomes 5)
	if items[0].Quality != expectedQuality {
		t.Errorf("Expected quality to be %d, got %d", expectedQuality, items[0].Quality)
	}
}
