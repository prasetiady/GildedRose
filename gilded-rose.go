package main

// import "fmt"

// type Item struct {
// 	name            string
// 	sellIn, quality int
// }
type Attribute struct {
	name            string
	sellIn, quality int
}

type Item interface {
	Update()
	// Name()
	// SellIn()
	// Quality()
}

type NormalItem struct {
	Attribute
}

func (item *NormalItem) Update() {
	if item.sellIn > 0 {
		item.sellIn -= 1
	}

	if item.sellIn > 0 {
		item.quality -= 1
	} else {
		item.quality -= 2
	}

	if item.quality < 0 {
		item.quality = 0
	}
}

type AgedBrieItem struct {
	Attribute
}

func (item *AgedBrieItem) Update() {
	if item.sellIn > 0 {
		item.sellIn -= 1
	}

	if item.quality < 50 {
		item.quality += 1
	}
}

type SulfurasItem struct {
	Attribute
}

func (item *SulfurasItem) Update() {}

type BackStagePassesItem struct {
	Attribute
}

func (item *BackStagePassesItem) Update() {
	if item.sellIn > 0 {
		item.sellIn -= 1
	}

	if item.sellIn > 10 {
		item.quality += 1
	} else if item.sellIn > 5 {
		item.quality += 2
	} else if item.sellIn > 0 {
		item.quality += 3
	} else {
		item.quality = 0
	}
}

type ConjuredItem struct {
	Attribute
}

func (item *ConjuredItem) Update() {
	if item.sellIn > 0 {
		item.sellIn -= 1
	}

	if item.sellIn > 0 {
		item.quality -= 2
	} else {
		item.quality -= 4
	}

	if item.quality < 0 {
		item.quality = 0
	}
}

func main() {}

// var items = []Item{
// 	Item{"+5 Dexterity Vest", 10, 20},
// 	Item{"Aged Brie", 2, 0},
// 	Item{"Elixir of the Mongoose", 5, 7},
// 	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
// 	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
// 	Item{"Conjured Mana Cake", 3, 6},
// }
//
// func main() {
// 	fmt.Println("OMGHAI!")
// 	fmt.Println("before")
// 	for _, i := range items {
// 		fmt.Println(i.name, i.sellIn, i.quality)
// 	}
//
// 	//fmt.Println("after")
// 	GlidedRose()
// 	fmt.Println("before")
// 	fmt.Println()
// 	for _, i := range items {
// 		fmt.Println(i.name, i.sellIn, i.quality)
// 	}
// }
//
// func GlidedRose() {
// 	for i := 0; i < len(items); i++ {
//
// 		if items[i].name != "Aged Brie" && items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
// 			if items[i].quality > 0 {
// 				if items[i].name != "Sulfuras, Hand of Ragnaros" {
// 					items[i].quality = items[i].quality - 1
// 				}
// 			}
// 		} else {
// 			if items[i].quality < 50 {
// 				items[i].quality = items[i].quality + 1
// 				if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
// 					if items[i].sellIn < 11 {
// 						if items[i].quality < 50 {
// 							items[i].quality = items[i].quality + 1
// 						}
// 					}
// 					if items[i].sellIn < 6 {
// 						if items[i].quality < 50 {
// 							items[i].quality = items[i].quality + 1
// 						}
// 					}
// 				}
// 			}
// 		}
//
// 		if items[i].name != "Sulfuras, Hand of Ragnaros" {
// 			items[i].sellIn = items[i].sellIn - 1
// 		}
//
// 		if items[i].sellIn < 0 {
// 			if items[i].name != "Aged Brie" {
// 				if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
// 					if items[i].quality > 0 {
// 						if items[i].name != "Sulfuras, Hand of Ragnaros" {
// 							items[i].quality = items[i].quality - 1
// 						}
// 					}
// 				} else {
// 					items[i].quality = items[i].quality - items[i].quality
// 				}
// 			} else {
// 				if items[i].quality < 50 {
// 					items[i].quality = items[i].quality + 1
// 				}
// 			}
// 		}
// 	}
//
// }
