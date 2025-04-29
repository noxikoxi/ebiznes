package utils

import (
	"fmt"
	"strconv"
)

func ValidateID(id string) (int, error) {
	validatedId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return validatedId, nil
}

func ValidateProductsFilters(minPriceStr, idxStr string) (float32, int, error) {
	minPrice := 0.0
	idx := 0
	var err, err1 error
	if minPriceStr != "" {
		minPrice, err = strconv.ParseFloat(minPriceStr, 32)
	}
	if idxStr != "" {
		idx, err1 = strconv.Atoi(idxStr)
	}
	if err != nil || err1 != nil {
		return 0, 0, fmt.Errorf("parse error: %v, %v", err, err1)
	}
	return float32(minPrice), idx, nil
}
