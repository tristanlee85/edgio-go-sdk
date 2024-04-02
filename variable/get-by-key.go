package variable

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
)

// GetByKey returns an environment variable by its key.
// Mandatory params:
// FilterParams.EnvID
// FilterParams.Key
// Returns the environment variable that matches the key or nil if no environment variables match the key.
func (c ClientStruct) GetByKey(params FilterParams) (common.Variable, error) {
	if params.EnvID == "" {
		return common.Variable{}, errors.New("'EnvID' is required")
	}

	if params.Key == "" {
		return common.Variable{}, errors.New("'Key' is required")
	}

	fullVariableList, err := c.List(params.EnvID)
	if err != nil {
		return common.Variable{}, errors.New(err.Error())
	}

	return utils.GetByAttr[common.Variable](
		common.FilterListParams[common.Variable]{Needle: params.Key, Haystack: fullVariableList.Items},
	), nil
}
