package slices

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Filter filtra un slice de acuerdo con la función callback especificada
func Filter[T constraints.Ordered](nums []T, callback func(T) bool) []T {
	result := make([]T, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	log(fmt.Sprintf("%s: the result is: %v", pkgName, result))
	return result
}
