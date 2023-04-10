package common

func Flat[T comparable](array [][]T) []T {
	flattenArray := []T{}
	for _, item := range array {
		flattenArray = append(flattenArray, item...)
	}
	return flattenArray
}
