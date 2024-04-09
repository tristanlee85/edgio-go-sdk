# edgio/env

This package groups Edgio Environment specific funcs.

## `env.NewClient(params common.ClientParams) (ClientStruct, error)`

```go
credentials := common.Creds{
  Key:    "some-api-key",
  Secret: "some-api-secret",
}

envClient, err := env.NewClient(common.ClientParams{
  Credentials: credentials,
  Config:      common.ClientConfig{AccessToken: "some-access-token"},
})
```

This func returns a new client with the provided parameters, with an access token already set, and the default edgio params values for service, scope and API version (which can be overwritten if needed) to interact with your application's environments.

### `env.NewClient` Mandatory Params

- `params.Credentials`
  - `params.Credentials.Key`: Edgio REST API Client Key. This is the API Client's identifier on Edgio. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.
  - `params.Credentials.Secret`: Edgio REST API Client Secret. This is the API Client' secret, unique for each API Client. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.

### `env.NewClient` Optional Params & Default Values

- `params.Credentials`
  - `params.Credentials.Scopes`: Edgio REST API Client scopes requested by the client. Different APIs needs different scopes. Refer to the [REST API docs](https://docs.edg.io/rest_api) to figure which ones you need.
    - Default value: `app.cache+app.cache.purge+app.waf+app.waf:edit+app.waf:read+app.accounts+app.config` (all scopes).
  - `params.Credentials.AuthURL`: Edgio REST API auth url. You will probably never need to change this, but we included the option just in case (e. g.: future enterprise self-hosted option).
    - Default value: `https://id.edgio.app/connect/token` (Edgio's default auth API url).
- `params.Config`
  - `params.Config.ApiVersion`: Intended REST API Version. Each one of the Edgio REST APIs has its own Version, that must be provided when creating the client.
    - Default value: `v0.1`
  - `params.Config.Service`: Intended REST API Service. Each one of the Edgio REST APIs has its own Service, that must be provided when creating the client.
    - Default value: `accounts`
  - `params.Config.Scope`: Intended REST API Scope. Each one of the Edgio REST APIs has its own Scope, that must be provided when creating the client.
    - Default value: `environments`
  - `params.Config.Url`: Edgio REST API resources url. You will probably never need to change this, but we included the option just in case (e. g.: future enterprise self-hosted option).
    - Default value: `https://edgioapis.com` (Edgio's default resources API url).

## `env.List(propertyID string) (ListResultType, error)`

```go
credentials := common.Creds{
  Key:    "some-api-key",
  Secret: "some-api-secret",
}

client, err := env.NewClient(env.ClientParams{
  Credentials: credentials,
  Config:      common.ClientConfig{},
})

envs, err := client.List("some-property-id") // [{ "ID": "some-id", "Name": "some-env-name", "LegacyAccNumber": "some-acc-number", "DefaultDomainName": "some-domain-name", "DNSDomainName": "some-dns", "CanMembersDeploy": true, "OnlyMaintainersCanDeploy": true, "HTTPRequestLogging": true, "PciCompliance": true, "CreatedAt": "2019-08-24T14:15:22Z", "UpdatedAt": "2019-08-24T14:15:22Z" }]
```

This func list environments for a given Edgio property. Edgio's list page size was defaulted to 100 for now, which is the highest value. The idea is to return all environments until actual pagination is implemented. Returns a list of environments for a given Property or an error if anything goes wrong.

### `env.List` Mandatory Params

- `propertyID`: Property ID from the property which owns the environments to be retrieved

### `env.List` Optional Params & Default Values

There is no optional parameters for that function

## `env.FilterList(params env.FilterParams) (common.FilteredListResultType[common.Env], error)`

```go
params := common.ClientParams{
  Credentials: common.Creds{ ... }
}

client, _ := env.NewClient(params)
filterParams := env.FilterParams{
  PropertyID: "some-property-id",
  Name:       "some-env",
}

FilteredList, _ := client.FilterList(filterParams)

fmt.Println(FilteredList) // [{ "ID": "some-id", "Name": "some-env", "LegacyAccNumber": "some-acc-number", "DefaultDomainName": "some-domain-name", "DNSDomainName": "some-dns", "CanMembersDeploy": true, "OnlyMaintainersCanDeploy": true, "HTTPRequestLogging": true, "PciCompliance": true, "CreatedAt": "2019-08-24T14:15:22Z", "UpdatedAt": "2019-08-24T14:15:22Z" }]
```

Filters the list of environments for a given Property by the environment name, and returns a list of environments for a given Property that contain the provided name, or all environments if no name is provided.

### `env.FilterList` Mandatory Params

- `env.FilterParams.propertyID`: Property ID from the property which owns the environments to be retrieved

### `env.FilterList` Optional Params

- `env.FilterParams.Name`: The string to be used as a environment name filter

## `env.Get(params env.FilterParams) (common.Env, error)`

```go
params := common.ClientParams{
  Credentials: common.Creds{ ... },
  Config: common.ClientConfig{OrgID: "some-org-id"},
}

client, _ := env.NewClient(params)
env, _ := client.Get(env.FilterParams{ ID: "some-env-id" })

fmt.Println(env) // { "ID": "some-env-id", "Name": "some-env", "LegacyAccNumber": "some-acc-number", "DefaultDomainName": "some-domain-name", "DNSDomainName": "some-dns", "CanMembersDeploy": true, "OnlyMaintainersCanDeploy": true, "HTTPRequestLogging": true, "PciCompliance": true, "CreatedAt": "2019-08-24T14:15:22Z", "UpdatedAt": "2019-08-24T14:15:22Z" }
```

This func retrieves an environment by ID and returns it, or empty if none was found.

### `env.Get` Mandatory Params

- `env.FilterParams.ID`: The string to be used as ID to get the desired environment

### `env.Get` Optional Params & Default Values

This func has no optional params.

## `env.GetByName(params FilterParams) (common.Env, error)`

```go

envClient, err := env.NewClient(common.ClientParams{
  Credentials: common.Creds{ ... },
})

env, _ := envClient.GetByName(env.FilterParams{PropertyID: "some-property-id", Name: "some-env-name"})

fmt.Println(env) // { "id": "some-id", "name": "some-env-name", "legacy_account_number": "", "default_domain_name": "", "dns_domain_name": "", "can_members_deploy": true, "only_maintainers_can_deploy": true, "http_request_logging": true, "pci_compliance": true, "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z" }
```

This func returns the first environment in the list that matches the name, or nil if no environments match the name.

### `env.GetByName` Mandatory Params

- `env.FilterParams.PropertyID`: Property ID from the property which owns the environments to be retrieved
- `env.FilterParams.Name`: The string to be used as a environment name filter

### `env.GetByName` Optional Params

There is no optional parameters for that function

<p align="right"><em><a href="../#edgioenvironment">back to the main README</a></em></p>
