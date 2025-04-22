package gildedrose

import (
	"fmt"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

type Updater interface {
	Update(item *Item)
}

type NormalItemUpdater struct{}

func (n NormalItemUpdater) Update(item *Item) {
	// decrease quality by 1 and then if sellIn < 0 decrease quality by 1 again
	// this simulates 'if sellIn has passed, quality degredes twice as fast'
	if item.Quality > 0 {
		item.Quality--
	}
	item.SellIn--
	if item.SellIn < 0 {
		if item.Quality > 0 {
			item.Quality--
		}
	}
}

type AgedBrieUpdater struct{}

func (a AgedBrieUpdater) Update(item *Item) {
	// increase quality by 1 and then if sellIn < 0 increase quality by 1 again
	// this simulates 'AgedBrie' increases in quality with time and
	// 'if sellIn has passed, quality increases twice as fast'
	if item.Quality < 50 {
		item.Quality++
	}
	item.SellIn--
	if item.SellIn < 0 {
		if item.Quality < 50 {
			item.Quality++
		}
	}
}

type SulfurasUpdater struct{}

func (s SulfurasUpdater) Update(item *Item) {
	// Sulfuras does not change in quality or sellIn
	// so we do nothing here
	// but we need to implement this interface to satisfy the Updater interface
	// so we can use it in the map
	item.SellIn--
	if item.SellIn < 0 {
		item.SellIn = 0
	}
}

type BackstagePassUpdater struct{}

func (b BackstagePassUpdater) Update(item *Item) {

	if item.Quality < 50 {
		if item.SellIn > 5 && item.SellIn <= 10 {
			item.Quality += 2
		} else if item.SellIn > 0 && item.SellIn <= 5 {
			item.Quality += 3
		} else if item.SellIn <= 0 {
			item.Quality = 0
		}
	}
	item.SellIn--
}

type ConjuredItemUpdater struct{}

func (c ConjuredItemUpdater) Update(item *Item) {
	// decrease quality by 2 and then if sellIn < 0 decrease quality by 2 again
	// this simulates 'conjured item degrades in quality twice as fast as normal'
	// and 'if sellIn has passed, quality degredes twice as fast'
	if item.Quality > 0 {
		item.Quality -= 2
	}
	item.SellIn--
	if item.SellIn < 0 {
		if item.Quality > 0 {
			item.Quality -= 2
		}
	}
}

var updaters = map[string]Updater{
	"Aged Brie":                                 AgedBrieUpdater{},
	"Sulfuras, Hand of Ragnaros":                SulfurasUpdater{},
	"Backstage passes to a TAFKAL80ETC concert": BackstagePassUpdater{},
	"Conjured":                                  ConjuredItemUpdater{},
	"Normal":                                    NormalItemUpdater{},
}

func UpdateQuality(items []*Item) {
	// composition over inheritance
	for _, item := range items {
		updater, ok := updaters[item.Name]
		if !ok {
			fmt.Printf("No updater found for item: %s\n", item.Name)
			fmt.Println("Defaulting to NormalItemUpdater")
			updater = NormalItemUpdater{}
		}
		updater.Update(item)
	}
}

func CreateSingleItem(name string, sellIn, quality int) []*Item {
	return []*Item{{Name: name, SellIn: sellIn, Quality: quality}}
}
