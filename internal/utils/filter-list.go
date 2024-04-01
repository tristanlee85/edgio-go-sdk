package utils

import (
	"edgio/common"
	"strings"
)

type Filterable interface {
	common.Searchable
	common.Property | common.Env | common.Variable
}

type FilterListParams[T Filterable] struct {
	Needle   string
	Haystack []T
}

// FilterList Filters the list of items by the given needle.
// Mandatory params:
// FilterListParams.Needle
// FilterListParams.Haystack
// Returns a list of items that contain the needle in their name, key or slug,
// depending on the entity type (Property, Environment, Variable),
// or an empty list if no items match the needle.
func FilterList[T Filterable](params FilterListParams[T]) []T {
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
