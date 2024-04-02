package utils

import (
	"edgio/common"
)

// GetByAttr returns the first item in the haystack that matches the needle.
// Mandatory params:
// common.FilterListParams.Needle
// common.FilterListParams.Haystack
// Returns the first item that matches the needle or nil if no items match the needle.
func GetByAttr[T common.Filterable](params common.FilterListParams[T]) T {
	var result T

	if params.Needle == "" || len(params.Haystack) == 0 {
		return result
	}

	for _, item := range params.Haystack {
		if item.GetKey() == params.Needle || item.GetName() == params.Needle || item.GetSlug() == params.Needle {
			result = item

			break
		}
	}

	return result
}
