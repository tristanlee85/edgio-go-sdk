package utils

import (
	"edgio/common"
)

func GetByAttr[T common.Filterable](params common.FilterListParams[T]) T {
	var result T

	for _, item := range params.Haystack {
		if item.GetKey() == params.Needle || item.GetName() == params.Needle || item.GetSlug() == params.Needle {
			result = item
			break
		}
	}

	return result
}
