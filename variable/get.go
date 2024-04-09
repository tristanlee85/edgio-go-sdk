package variable

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
	"net/http"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

// Get retrieves an environment variable by ID
// mandatory params: FilterParams.ID
// returns retrieved environment variable, or empty if none was found.
func (c ClientStruct) Get(params FilterParams) (common.Variable, error) {
	var variableGetResult common.Variable

	if params.ID == "" {
		return common.Variable{}, errors.New("'ID' is required")
	}

	httpClient := &http.Client{}
	serviceURL := c.GetServiceURL(common.URLParams{Path: params.ID})

	parsedURL, err := url.Parse(serviceURL)
	if err != nil {
		return common.Variable{}, errors.New(err.Error())
	}

	request, _ := http.NewRequest(http.MethodGet, parsedURL.String(), nil)

	variableJSONmap, err := utils.GetHTTPJSONResult(httpClient, request, c.AccessToken)
	if err != nil {
		return common.Variable{}, errors.New(err.Error())
	}

	err = mapstructure.Decode(variableJSONmap, &variableGetResult)
	if err != nil {
		return common.Variable{}, errors.New(err.Error())
	}

	return variableGetResult, nil
}
