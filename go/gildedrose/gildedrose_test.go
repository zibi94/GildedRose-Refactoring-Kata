package gildedrose

import (
	"testing"
)

func TestHandlers(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		item := &Item{"Foo", 5, 10}

		for i := 0; i < 10; i++ {
			defaultHandler(item)
			item.checkPoint(t, 1, 6)
			item.checkPoint(t, 0, 4)
		}

		item.requireQuality(t, 0)
	})

	t.Run(agedBrie, func(t *testing.T) {
		item := &Item{agedBrie, 5, 10}

		for i := 0; i < 5; i++ {
			agedBrieHandler(item)
		}

		item.requireQuality(t, 15)
	})

	t.Run(backstagePasses, func(t *testing.T) {
		item := &Item{backstagePasses, 13, 10}

		for i := 0; i < 13; i++ {
			backstagePassesHandler(item)
			item.checkPoint(t, 11, 12)
			item.checkPoint(t, 10, 14)
			item.checkPoint(t, 5, 25)
		}

		item.requireQuality(t, 0)
	})

	t.Run(sulfuras, func(t *testing.T) {
		item := &Item{sulfuras, 10, 25}

		for i := 0; i < 15; i++ {
			sulfurasHandler(item)
			item.checkPoint(t, 4, 25)
		}

		item.requireQuality(t, 25)
	})

	t.Run(conjured, func(t *testing.T) {
		item := &Item{conjured, 10, 20}

		for i := 0; i < 10; i++ {
			conjuredHandler(item)
			item.checkPoint(t, 5, 10)
		}

		item.requireQuality(t, 0)
	})
}

func TestItemMethods(t *testing.T) {
	t.Run("Align_Quality", func(t *testing.T) {
		item := &Item{conjured, 10, -2}

		item.alignQuality()
		item.requireQuality(t, 0)

		item.Quality = 52

		item.alignQuality()
		item.requireQuality(t, 50)
	})
}

func (i *Item) checkPoint(t *testing.T, day, expectQlt int) {
	if i.SellIn == day && i.Quality != expectQlt {
		t.Errorf("Day: %d. Quality: Expected %d but got %d.", i.SellIn, expectQlt, i.Quality)
	}
}

func (i *Item) requireQuality(t *testing.T, expected int) {
	if i.Quality != expected {
		t.Errorf("Quality: Expected %d but got %d.", expected, i.Quality)
	}
}
