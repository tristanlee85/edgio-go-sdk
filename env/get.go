package env

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
	"net/http"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

// Get retrieves an env by ID
// mandatory params: FilterParams.ID
// returns retrieved environment, or empty if none was found.
func (c ClientStruct) Get(params FilterParams) (common.Env, error) {
	var envGetResult common.Env

	if params.ID == "" {
		return common.Env{}, errors.New("'ID' is required")
	}

	httpClient := &http.Client{}
	serviceURL := c.GetServiceURL(common.URLParams{Path: params.ID})

	parsedURL, err := url.Parse(serviceURL)
	if err != nil {
		return common.Env{}, errors.New(err.Error())
	}

	request, _ := http.NewRequest(http.MethodGet, parsedURL.String(), nil)

	envJSONmap, err := utils.GetHTTPJSONResult(httpClient, request, c.AccessToken)
	if err != nil {
		return common.Env{}, errors.New(err.Error())
	}

	err = mapstructure.Decode(envJSONmap, &envGetResult)
	if err != nil {
		return common.Env{}, errors.New(err.Error())
	}

	return envGetResult, nil
}
