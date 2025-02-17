# edgio/variables

This package groups Edgio Environment Variables specific funcs.

## `variable.NewClient(params common.ClientParams) (ClientStruct, error)`

```go
credentials := common.Creds{
  Key:    "some-api-key",
  Secret: "some-api-secret",
}

envClient, err := variable.NewClient(common.ClientParams{
  Credentials: credentials,
  Config:      common.ClientConfig{AccessToken: "some-access-token"},
})
```

This func returns a new client with the provided parameters, with an access token already set, and the default edgio params values for service, scope and API version (which can be overwritten if needed) to interact with your application's environments.

### `variable.NewClient` Mandatory Params

- `params.Credentials`
  - `params.Credentials.Key`: Edgio REST API Client Key. This is the API Client's identifier on Edgio. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.
  - `params.Credentials.Secret`: Edgio REST API Client Secret. This is the API Client' secret, unique for each API Client. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.

### `variable.NewClient` Optional Params & Default Values

- `params.Credentials`
  - `params.Credentials.Scopes`: Edgio REST API Client scopes requested by the client. Different APIs needs different scopes. Refer to the [REST API docs](https://docs.edg.io/rest_api) to figure which ones you need.
    - Default value: `app.cache+app.cache.purge+app.waf+app.waf:edit+app.waf:read+app.accounts+app.config` (all scopes).
  - `params.Credentials.AuthURL`: Edgio REST API auth url. You will probably never need to change this, but we included the option just in case (e. g.: future enterprise self-hosted option).
    - Default value: `https://id.edgio.app/connect/token` (Edgio's default auth API url).
- `params.Config`
  - `params.Config.ApiVersion`: Intended REST API Version. Each one of the Edgio REST APIs has its own Version, that must be provided when creating the client.
    - Default value: `v0.1`
  - `params.Config.Service`: Intended REST API Service. Each one of the Edgio REST APIs has its own Service, that must be provided when creating the client.
    - Default value: `config`
  - `params.Config.Scope`: Intended REST API Scope. Each one of the Edgio REST APIs has its own Scope, that must be provided when creating the client.
    - Default value: `environment-variables`
  - `params.Config.Url`: Edgio REST API resources url. You will probably never need to change this, but we included the option just in case (e. g.: future enterprise self-hosted option).
    - Default value: `https://edgioapis.com` (Edgio's default resources API url).

## `variable.List(environmentID string) (ListResultType, error)`

```go
credentials := common.Creds{
  Key:    "some-api-key",
  Secret: "some-api-secret",
}

client, err := variable.NewClient(variable.ClientParams{
  Credentials: credentials,
  Config:      common.ClientConfig{},
})

variables, err := client.List("some-environment-id") // [{ "ID": "string", "Key": "string", "Value": "string", "Secret": true, "CreatedAt": "2019-08-24T14:15:22Z", "UpdatedAt": "2019-08-24T14:15:22Z" }]
```

This func list environment variables for a given Edgio Environment. Edgio's list page size was defaulted to 100 for now, which is the highest value. The idea is to return all environment variables until actual pagination is implemented. Returns a list of environment variables for a given Environment or an error if anything goes wrong.

### `variable.List` Mandatory Params

- `environmentID`: Property ID from the property which owns the environments to be retrieved

### `variable.List` Optional Params & Default Values

There is no optional parameters for that function

## `variable.FilterList(params variable.FilterParams) (common.FilteredListResultType[common.Property], error)`

```go
params := common.ClientParams{
  Credentials: common.Creds{ ... },
  Config: common.ClientConfig{OrgID: "some-org-id"},
}

client, _ := variable.NewClient(params)
List, _ := client.FilterList(variable.FilterParams{ Key: "some-Key" }) // [{ "ID": "string", "Key": "some-key", "Value": "some-value", "Secret": true, "CreatedAt": "2019-08-24T14:15:22Z", "UpdatedAt": "2019-08-24T14:15:22Z" }]
```

This func filters the list of environment variables for a given Environment by the variable key, and returns a list of environment variables that contain the provided key, or all environment variables if no key is provided.

### `variable.FilterList` Mandatory Params

- `variable.FilterParams.EnvID`: The environment ID to get variables from

### `variable.FilterList` Optional Params & Default Values

- `variable.FilterParams.Key`: The string to be used as parameter for the list filter

## `variable.Get(params variable.FilterParams) (common.Variable, error)`

```go
params := common.ClientParams{
  Credentials: common.Creds{ ... },
  Config: common.ClientConfig{OrgID: "some-org-id"},
}

client, _ := variable.NewClient(params)
variable, _ := client.Get(variable.FilterParams{ ID: "some-variable-id" })

fmt.Println(variable) // { "ID": "some-variable-id", "Key": "some-key", "Value": "some-value", "Secret": true, "CreatedAt": "2019-08-24T14:15:22Z", "UpdatedAt": "2019-08-24T14:15:22Z" }
```

This func retrieves an environment variable by ID and returns it, or empty if none was found.

### `variable.Get` Mandatory Params

- `variable.FilterParams.ID`: The string to be used as ID to get the desired environment variable

### `variable.Get` Optional Params & Default Values

This func has no optional params.

## `GetByKey(params FilterParams) (common.Variable, error)`

```go
variableClient, _ := variable.NewClient(common.ClientParams{
  Credentials: common.Creds{ ... },
})

variable, _ := variableClient.GetByKey(variable.FilterParams{EnvID: "some-env-id", Key: "SOME_ENV_VAR"})

fmt.Println(variable) // "id": "some-id", "key": "some-env-var-key", "value": "some-value", "secret": true, "create_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z" }
```

This func returns the environment variable that matches the provided key, or nil if no environment variables match the key.

### `variable.GetByKey` Mandatory Params

- `variable.FilterParams.EnvID`: The environment ID to get variables from
- `variable.FilterParams.Key`: The string to be used as parameter for the list filter

### `variable.GetByKey` Optional Params & Default Values

There is no optional parameters for that function

<p align="right"><em><a href="../#edgiovariables">back to the main README</a></em></p>
