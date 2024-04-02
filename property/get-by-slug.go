package property

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
)

// GetBySlug returns the first property in the list that matches the slug.
// Mandatory params:
// FilterParams.Slug
// Returns the first property that matches the slug or nil if no properties match the slug.
func (c ClientStruct) GetBySlug(params FilterParams) (common.Property, error) {
	fullPropertyList, err := c.List()
	if err != nil {
		return common.Property{}, errors.New(err.Error())
	}

	if params.Slug == "" {
		return common.Property{}, errors.New("'Slug' is required")
	}

	return utils.GetByAttr[common.Property](
		common.FilterListParams[common.Property]{Needle: params.Slug, Haystack: fullPropertyList.Items},
	), nil
}
