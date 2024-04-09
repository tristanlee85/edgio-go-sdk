package property

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
	"net/http"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

// Get retrieves a property by ID
// mandatory params: FilterParams.ID
// returns retrieved property, or empty if none was found.
func (c ClientStruct) Get(params FilterParams) (common.Property, error) {
	var propertyGetResult common.Property

	if params.ID == "" {
		return common.Property{}, errors.New("'ID' is required")
	}

	httpClient := &http.Client{}
	serviceURL := c.GetServiceURL(common.URLParams{Path: params.ID})

	parsedURL, err := url.Parse(serviceURL)
	if err != nil {
		return common.Property{}, errors.New(err.Error())
	}

	request, _ := http.NewRequest(http.MethodGet, parsedURL.String(), nil)

	propertiesJSONmap, err := utils.GetHTTPJSONResult(httpClient, request, c.AccessToken)
	if err != nil {
		return common.Property{}, errors.New(err.Error())
	}

	err = mapstructure.Decode(propertiesJSONmap, &propertyGetResult)
	if err != nil {
		return common.Property{}, errors.New(err.Error())
	}

	return propertyGetResult, nil
}
