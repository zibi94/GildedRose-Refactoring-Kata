package gildedrose

// Default handler.
func defaultHandler(i *Item) {
	i.SellIn--
	if i.Quality == 0 {
		return
	}

	if i.SellIn <= 0 {
		i.Quality--
	}

	i.Quality--
	i.alignQuality()
}

// Custom handlers.
var handlers = map[string]func(i *Item){
	agedBrie:        agedBrieHandler,
	backstagePasses: backstagePassesHandler,
	sulfuras:        sulfurasHandler,
	conjured:        conjuredHandler,
}

func agedBrieHandler(i *Item) {
	i.SellIn--
	i.Quality++
	i.alignQuality()
}

func backstagePassesHandler(i *Item) {
	i.SellIn--
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
}

func sulfurasHandler(i *Item) {
	// Legendary item - never has to be sold or decreases in Quality.
}

func conjuredHandler(i *Item) {
	i.SellIn--
	if i.Quality == 0 {
		return
	}

	i.Quality -= 2
	i.alignQuality()
}
