package common

type FilterFunc[T comparable] func(T) bool

func Filter[T comparable](array []T, filterFunc FilterFunc[T]) []T {
	filteredArray := []T{}
	for _, item := range array {
		if filterFunc(item) {
			filteredArray = append(filteredArray, item)
		}
	}
	return filteredArray
}
