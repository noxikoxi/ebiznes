package scopes

import (
	"gorm.io/gorm"
)

// GORM SCOPE
func ProductCategory(categoryID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("category_id = ?", categoryID)
	}
}

func PriceAbove(price float32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price > ?", price)
	}
}

func CartID(cartID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("cart_id = ?", cartID)
	}
}

func CartProductID(productID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("product_id = ?", productID)
	}
}

func Name(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", name)
	}
}

func GetCategoryPriceFilters(price float32, categoryID uint) []func(db *gorm.DB) *gorm.DB {
	var filters []func(db *gorm.DB) *gorm.DB

	if price != 0 {
		filters = append(filters, PriceAbove(price)) // Dodajemy filtr ceny do listy
	}

	if categoryID != 0 {
		filters = append(filters, ProductCategory(categoryID)) // Dodajemy filtr kategorii do listy
	}

	return filters
}
