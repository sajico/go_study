package common

import "golang.org/x/exp/constraints"

type FindFunc[T comparable] func(T, T) T

func find[T comparable](array []T, findFunc FindFunc[T]) T {
	findItem := array[0]
	for _, item := range array {
		findItem = findFunc(findItem, item)
	}
	return findItem
}

func lessthan[T constraints.Ordered](a T, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func greaterthan[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min[T constraints.Ordered](array []T) T {
	return find(
		array,
		func(a T, b T) T { return lessthan(a, b) },
	)
}

func Max[T constraints.Ordered](array []T) T {
	return find(
		array,
		func(a T, b T) T { return greaterthan(a, b) },
	)
}
