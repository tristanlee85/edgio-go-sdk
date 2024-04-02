package env

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
)

// GetByName returns the first environment in the list that matches the name.
// Mandatory params:
// FilterParams.PropertyID
// FilterParams.Name
// Returns the first environment that matches the name or nil if no environments match the name.
func (c ClientStruct) GetByName(params FilterParams) (common.Env, error) {
	if params.PropertyID == "" {
		return common.Env{}, errors.New("'PropertyID' is required")
	}

	fullEnvList, err := c.List(params.PropertyID)
	if err != nil {
		return common.Env{}, errors.New(err.Error())
	}

	if params.PropertyID == "" {
		return common.Env{}, errors.New("'Name' is required")
	}

	return utils.GetByAttr[common.Env](
		common.FilterListParams[common.Env]{Needle: params.Name, Haystack: fullEnvList.Items},
	), nil
}
