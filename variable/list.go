package variable

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
	"net/http"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

type ListResult common.BaseListResultType[common.Variable]

var variableListResult ListResult

// List Lists the environment variables for a given Environment.
// Edgio's list page size was defaulted to 100 for now,
// which is the highest value. The idea is to return all environment variables
// until actual pagination is implemented.
// Returns a list of environment variables for a given Environment, or an error if anything goes wrong.
func (c ClientStruct) List(environmentID string) (ListResult, error) {
	httpClient := &http.Client{}
	serviceURL := c.GetServiceURL(common.URLParams{})

	parsedURL, err := url.Parse(serviceURL)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	rawQueryString := utils.ToQueryString(
		map[string]string{"page_size": "100", "environment_id": environmentID},
	)

	parsedURL.RawQuery = rawQueryString

	request, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	variableJSONmap, err := utils.GetHTTPJSONResult(httpClient, request, c.AccessToken)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	err = mapstructure.Decode(variableJSONmap, &variableListResult)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	return variableListResult, nil
}

// FilterList Filters the list of environment variables for a given Environment by the variable key.
// Mandatory params:
// variable.FilterParams.EnvID
// Optional params:
// variable.FilterParams.Key
// Returns a list of environment variables that contain the provided key,
// or all environment variables if no key is provided.
func (c ClientStruct) FilterList(params FilterParams) (common.FilteredListResultType[common.Variable], error) {
	if params.EnvID == "" {
		return common.FilteredListResultType[common.Variable]{}, errors.New("EnvID is required")
	}

	fullVarList, err := c.List(params.EnvID)
	if err != nil {
		return common.FilteredListResultType[common.Variable]{}, errors.New(err.Error())
	}

	if params.Key == "" {
		return common.FilteredListResultType[common.Variable]{
			BaseListResultType: common.BaseListResultType[common.Variable]{
				Total: fullVarList.Total,
				Items: fullVarList.Items,
			},
			FilteredTotal: fullVarList.Total,
		}, nil
	}

	filteredProperties := utils.FilterList[common.Variable](
		common.FilterListParams[common.Variable]{Needle: params.Key, Haystack: fullVarList.Items},
	)

	return common.FilteredListResultType[common.Variable]{
		BaseListResultType: common.BaseListResultType[common.Variable]{
			Total: fullVarList.Total,
			Items: filteredProperties,
		},
		FilteredTotal: len(filteredProperties),
	}, nil
}
