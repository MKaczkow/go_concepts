package gildedrose_test

import (
	"gilded_rose/gildedrose"
	"testing"
)

func Test_UpdateQuality_BackstagePass(t *testing.T) {
	tests := []struct {
		name            string
		initialSellIn   int
		initialQuality  int
		expectedSellIn  int
		expectedQuality int
	}{
		{"5 days left", 5, 10, 4, 13},
		{"10 days left", 10, 10, 9, 12},
		{"more than 10 days left", 15, 10, 14, 10},
		{"on sell-by date", 0, 10, -1, 0},
		{"one day after sell-by", -1, 10, -2, 0},
		{"quality at max", 5, 50, 4, 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			items := gildedrose.CreateSingleItem("Backstage passes to a TAFKAL80ETC concert", test.initialSellIn, test.initialQuality)
			gildedrose.UpdateQuality(items)
			if items[0].SellIn != test.expectedSellIn {
				t.Errorf("Expected SellIn to be %d, got %d", test.expectedSellIn, items[0].SellIn)
			}
			if items[0].Quality != test.expectedQuality {
				t.Errorf("Expected Quality to be %d, got %d", test.expectedQuality, items[0].Quality)
			}
		})
	}
}
