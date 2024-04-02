package utils_test

import (
	"edgio/common"
	"edgio/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByAttr(t *testing.T) {
	propertyHaystack := []common.Property{
		{Slug: "slug1"},
		{Slug: "slug2"},
		{Slug: "slug3"},
	}

	envHaystack := []common.Env{
		{Name: "name1"},
		{Name: "name2"},
		{Name: "name3"},
	}

	variableHaystack := []common.Variable{
		{Key: "key1"},
		{Key: "key2"},
		{Key: "key3"},
	}

	t.Run("should return item that match the needle (with 'Slug')", func(t *testing.T) {
		params := common.FilterListParams[common.Property]{Needle: "slug1", Haystack: propertyHaystack}
		result := utils.GetByAttr[common.Property](params)

		assert.Equal(t, "slug1", result.GetSlug())
	})

	t.Run("should return item that match the needle (with 'Name')", func(t *testing.T) {
		params := common.FilterListParams[common.Env]{Needle: "name1", Haystack: envHaystack}
		result := utils.GetByAttr[common.Env](params)

		assert.Equal(t, "name1", result.GetName())
	})

	t.Run("should return item that match the needle (with 'Key')", func(t *testing.T) {
		params := common.FilterListParams[common.Variable]{Needle: "key1", Haystack: variableHaystack}
		result := utils.GetByAttr[common.Variable](params)

		assert.Equal(t, "key1", result.GetKey())
	})

	t.Run("should return no items when none match the needle", func(t *testing.T) {
		params := common.FilterListParams[common.Property]{Needle: "slug4", Haystack: propertyHaystack}
		result := utils.GetByAttr[common.Property](params)

		assert.Empty(t, result)
	})

	t.Run("should match only if needle is equal attr (return empty)", func(t *testing.T) {
		params := common.FilterListParams[common.Property]{Needle: "lug1", Haystack: propertyHaystack}
		result := utils.GetByAttr[common.Property](params)

		assert.Empty(t, result)
	})

	t.Run("should return empty if no needle is provided", func(t *testing.T) {
		params := common.FilterListParams[common.Property]{Haystack: propertyHaystack}
		result := utils.GetByAttr[common.Property](params)

		assert.Empty(t, result)
	})

	t.Run("should return empty if no haystack is provided", func(t *testing.T) {
		params := common.FilterListParams[common.Property]{Needle: "slug1"}
		result := utils.GetByAttr[common.Property](params)

		assert.Empty(t, result)
	})
}
