package scopes

import (
	"gorm.io/gorm"
	"strconv"
)

// GORM SCOPE
func ProductCategory(categoryID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("category_id = ?", categoryID)
	}
}

func PriceAbove(price float64) func(db *gorm.DB) *gorm.DB {
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

func GetCategoryPriceFilters(price string, categoryID string) []func(db *gorm.DB) *gorm.DB {
	var filters []func(db *gorm.DB) *gorm.DB

	if price != "" {
		minPrice, err := strconv.ParseFloat(price, 64)
		if err == nil {
			filters = append(filters, PriceAbove(minPrice)) // Dodajemy filtr ceny do listy
		}
	}

	// Obsługa filtra kategorii
	if categoryID != "" {
		id, err := strconv.Atoi(categoryID)
		if err == nil {
			filters = append(filters, ProductCategory(uint(id))) // Dodajemy filtr kategorii do listy
		}
	}

	return filters
}
