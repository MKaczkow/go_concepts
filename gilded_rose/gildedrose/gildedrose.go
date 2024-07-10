package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const (
	AgedBrie             = "Aged Brie"
	Sulfuras             = "Sulfuras, Hand of Ragnaros"
	BackstagePasses      = "Backstage passes to a TAFKAL80ETC concert"
	Conjured             = "Conjured"
	MaxQuality           = 50
	LegendaryItemQuality = 80
)

func UpdateQuality(items []*Item) {
	for _, item := range items {
		switch item.Name {
		case AgedBrie:
			updateAgedBrie(item)
		case Sulfuras:
			// Legendary item, no need to update quality or sell-in
			continue
		case BackstagePasses:
			updateBackstagePasses(item)
		case Conjured:
			updateConjured(item)
		default:
			updateNormalItem(item)
		}

		// Ensure quality is within bounds
		if item.Quality > MaxQuality {
			item.Quality = MaxQuality
		} else if item.Quality < 0 {
			item.Quality = 0
		}

		// Decrease sell-in for all items except legendary ones
		if item.Name != Sulfuras {
			item.SellIn--
		}

		// Adjust quality when sell-in date has passed
		if item.SellIn < 0 {
			handleExpiredItem(item)
		}
	}
}

func updateAgedBrie(item *Item) {
	if item.Quality < MaxQuality {
		item.Quality++
	}
}

func updateBackstagePasses(item *Item) {
	if item.Quality < MaxQuality {
		item.Quality++
		if item.SellIn < 11 {
			if item.Quality < MaxQuality {
				item.Quality++
			}
		}
		if item.SellIn < 6 {
			if item.Quality < MaxQuality {
				item.Quality++
			}
		}
	}
	if item.SellIn < 0 {
		item.Quality = 0
	}
}

func updateConjured(item *Item) {
	if item.Quality > 0 {
		item.Quality -= 2
	}
}

func updateNormalItem(item *Item) {
	if item.Quality > 0 {
		item.Quality--
		if item.SellIn < 0 {
			item.Quality--
		}
	}
}

func handleExpiredItem(item *Item) {
	switch item.Name {
	case AgedBrie:
		updateAgedBrie(item)
	case BackstagePasses:
		item.Quality = 0
	default:
		updateNormalItem(item)
	}
}
