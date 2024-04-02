package utils

import (
	"edgio/common"
	"strings"
)

// FilterList Filters the list of items by the given needle.
// Mandatory params:
// common.FilterListParams.Needle
// common.FilterListParams.Haystack
// Returns a list of items that contain the needle in their name, key or slug,
// depending on the entity type (Property, Environment, Variable),
// or an empty list if no items match the needle.
func FilterList[T common.Filterable](params common.FilterListParams[T]) []T {
	result := []T{}

	for _, item := range params.Haystack {
		if strings.Contains(item.GetKey(), params.Needle) ||
			strings.Contains(item.GetName(), params.Needle) ||
			strings.Contains(item.GetSlug(), params.Needle) {
			result = append(result, item)
		}
	}

	return result
}
