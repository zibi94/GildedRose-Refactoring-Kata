package gildedrose

const (
	// Quality extremes.
	maxQuality = 50
	minQuality = 0

	// Items which need custom quality handling.
	agedBrie        string = "Aged Brie"
	backstagePasses string = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        string = "Sulfuras, Hand of Ragnaros"
	conjured        string = "Conjured Mana Cake" // New item.
)

type Item struct {
	Name            string
	SellIn, Quality int
}

// Handle choose handler by name and execute.
// If there is no custom handler for item use default handler.
func (i *Item) handle() {
	if fn, ok := handlers[i.Name]; ok {
		fn(i)
	} else {
		defaultHandler(i)
	}
}

// The method sets the minimum and maximum quality if the value exceeds the extreme values.
func (i *Item) alignQuality() {
	if i.Quality < minQuality {
		i.Quality = minQuality
	}
	if i.Quality > maxQuality {
		i.Quality = maxQuality
	}
}

// UpdateQuality decrease expiration date and quality.
// The function should be used at the end of the day because it lowers the expiration date
// and recalculates the quality to the next day.
// Also func support custom handling for items:
// - "Aged Brie";
// - "Backstage passes to a TAFKAL80ETC concert";
// - "Sulfuras, Hand of Ragnaros"
// - "Conjured Mana Cake".
func UpdateQuality(items []*Item) {
	for i := range items {
		items[i].handle()
	}
}
