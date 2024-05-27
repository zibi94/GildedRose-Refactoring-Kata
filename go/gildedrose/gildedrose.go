package gildedrose

const (
	// Quality extremes.
	MaxQuality = 50
	MinQuality = 0

	// Items which need custom quality handling.
	AgedBrie        = "Aged Brie"
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	Sulfuras        = "Sulfuras, Hand of Ragnaros"
	Conjured        = "Conjured Mana Cake" // New item.
)

type Item struct {
	Name            string
	SellIn, Quality int
}

// The method sets the minimum and maximum quality if the value exceeds the extreme values.
func (i *Item) alignQuality() {
	if i.Quality < MinQuality {
		i.Quality = MinQuality
	}
	if i.Quality > MaxQuality {
		i.Quality = MaxQuality
	}
}

func UpdateQuality(items []*Item) {
	for i := range items {
		if fn, ok := handlers[items[i].Name]; ok {
			fn(items[i])
		} else {
			defaultHandler(items[i])
		}
	}

}

// Custom handlers.
var handlers = map[string]func(i *Item){
	AgedBrie: func(i *Item) {
		i.SellIn--
		i.Quality++
		i.alignQuality()
	},

	BackstagePasses: func(i *Item) {
		i.SellIn -= 1
		if i.Quality == 0 {
			return
		}

		if i.SellIn <= 0 {
			i.Quality = 0
			return
		}

		if i.SellIn <= 10 {
			i.Quality++
		}
		if i.SellIn <= 5 {
			i.Quality++
		}

		i.Quality++
		i.alignQuality()
	},

	Sulfuras: func(i *Item) {
		// Legendary item - never has to be sold or decreases in Quality.
	},

	Conjured: func(i *Item) {
		i.SellIn--
		if i.Quality == 0 {
			return
		}

		i.Quality -= 2
		i.alignQuality()
	},
}

// Default handler.
func defaultHandler(i *Item) {
	// Fix foo item name.
	if i.Name == "foo" {
		i.Name = "fixme"
	}

	i.SellIn -= 1
	if i.Quality == 0 {
		return
	}

	if i.SellIn <= 0 {
		i.Quality--
	}

	i.Quality--
	i.alignQuality()
}
