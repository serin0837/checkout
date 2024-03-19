package models

// SKU type is to keep track of Stock Keeping Units.
// type SKU struct {
// 	Items []Item
// }

// Item represent field for one item.
type Item struct {
	Name         string
	UnitPrice    int
	SpecialPrice int
	Number       int
}
