# internal/utils

This package package holds some utility functions that are used to outsource some common logic from other packages to avoid repetition/ease testing.

## `utils.GetHttpJsonResult(httpClient *http.Client, request *http.Request, token string) (map[string]any, error)`

This function has mainly three related goals:

1. Process http requests;
2. treat HTTP errors in a standardized way, and;
3. Process and decode returned json data from the endpoints.

```go
type resultType struct { Id string }
var result = resultType{ Id: `json:"id"` }

httpClient := &http.Client{}
request, _ := http.NewRequest(http.MethodGet, client.GetServiceURL(common.URLParams{ Path: "some-org-id" }), nil)

result, err := utils.GetHttpJsonResult(httpClient, request, "some-access-token")
```

### `utils.GetHttpJsonResult` Mandatory Params

- `httpClient *http.Client`: Http client instance from the `net/http` package.
- `request *http.Request`: The request configs (result of `net/http.NewRequest` function).
- `token string`: The Edgio API access token.

### `utils.GetHttpJsonResult` Optional Params & Default Values

There is no optional parameters for that function

<p align="right"><em><a href="../../#utils">back to the main README</a></em></p>

## `utils.FilterList[T common.Filterable](params common.FilterListParams[T]) []T`

Filters the list of items (haystack) by the given needle. Returns a list of items that contain the needle in their name, key or slug, depending on the entity type (Property, Environment, Variable), or an empty list if no items match the needle.

```go
haystack := []common.Variable{
  {Key: "nope"},
  {Key: "not-this"},
  {Key: "this-variable"},
}

filteredProperties := utils.FilterList[common.Variable](
  common.FilterListParams[common.Variable]{Needle: "this-variable", Haystack: haystack},
)

fmt.Println(filteredProperties) // [{ Key: "this-variable" }]
```

### `utils.FilterList` Mandatory Params

- `common.FilterListParams.Needle`: The string that should be used to filter the provided list
- `common.FilterListParams.Haystack`:  The list of items to be filtered byt the needle

### `utils.FilterList` Optional Params

This func has no optional params.

<p align="right"><em><a href="../../#utils">back to the main README</a></em></p>
