package utils

func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)

	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func Map[T, K any](slice []T, fn func(T) K) []K {
	result := make([]K, len(slice))

	for i, item := range slice {
		result[i] = fn(item)
	}

	return result
}
