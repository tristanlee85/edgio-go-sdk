# edgio/property

This package groups Edgio Property specific funcs.

## `property.NewClient(params common.ClientParams) (ClientStruct, error)`

```go
credentials := common.Creds{
  Key:    "some-api-key",
  Secret: "some-api-secret",
}

propertyClient, err := property.NewClient(common.ClientParams{
  Credentials: credentials,
  Config:      common.ClientConfig{OrgID: "some-org-id", AccessToken: "some-access-token"},
})
```

This func returns a new client with the provided parameters, with an access token already set, and the default edgio params values for service, scope and API version (which can be overwritten if needed) to interact with your application's properties.

### `property.NewClient` Mandatory Params

- `params.Credentials`
  - `params.Credentials.Key`: Edgio REST API Client Key. This is the API Client's identifier on Edgio. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.
  - `params.Credentials.Secret`: Edgio REST API Client Secret. This is the API Client' secret, unique for each API Client. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.
- `params.Config`
  - `params.Config.OrdID`: Edgio's OrganizationID where the properties-to-be-accessed lives under.

### `property.NewClient` Optional Params & Default Values

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
    - Default value: `properties`
  - `params.Config.Url`: Edgio REST API resources url. You will probably never need to change this, but we included the option just in case (e. g.: future enterprise self-hosted option).
    - Default value: `https://edgioapis.com` (Edgio's default resources API url).

## `property.List() (ListResultType, error)`

```go
credentials := common.Creds{
  Key:    "some-api-key",
  Secret: "some-api-secret",
}

client, err := property.NewClient(common.ClientParams{
  Credentials: credentials,
  Config:      common.ClientConfig{OrgID: "some-org-id"},
})

properties, err := client.List(common.URLParams{Path: "some-org-id"}) // [{ "ID": "prop-id", "Slug": "prop-slug", "CreatedAt": "2019-08-24T14:15:22Z", "UpdatedAt": "2019-08-24T14:15:22Z" }]
```

This func lists properties for a given Edgio Organization. Edgio's list page size was defaulted to 100 for now, which is the highest value. The idea is to return all properties until actual pagination is implemented. Returns a list of properties for a given Organization or an error if anything goes wrong.

### `property.List` Mandatory Params

There is no mandatory parameters for that function

### `property.List` Optional Params & Default Values

There is no optional parameters for that function

## `property.FilterList(params property.FilterParams) (common.FilteredListResultType[common.Property], error)`

```go
params := common.ClientParams{
  Credentials: common.Creds{ ... },
  Config: common.ClientConfig{OrgID: "some-org-id"},
}

client, _ := property.NewClient(params)
list, _ := client.FilterList(property.FilterParams{ Slug: "some-slug" })

fmt.Println(list) // [{ "ID": "prop-id", "Slug": "some-slug", "CreatedAt": "2019-08-24T14:15:22Z", "UpdatedAt": "2019-08-24T14:15:22Z" }]
```

Filters the list of properties for a given Org by the property slug, and returns a list of properties that contain the provided slug, or all properties if no slug is provided.

### `property.FilterList` Mandatory Params

This func has no mandatory params.

### `property.FilterList` Optional Params & Default Values

- `property.FilterParams.Slug`: The string to be used as slug to filter the property list

## `property.Get(params property.FilterParams) (common.Property, error)`

```go
params := common.ClientParams{
  Credentials: common.Creds{ ... },
  Config: common.ClientConfig{OrgID: "some-org-id"},
}

client, _ := property.NewClient(params)
property, _ := client.Get(property.FilterParams{ ID: "some-property-id" })

fmt.Println(property) // { "ID": "some-property-id", "Slug": "some-slug", "CreatedAt": "2019-08-24T14:15:22Z", "UpdatedAt": "2019-08-24T14:15:22Z" }
```

This func retrieves a property by ID and returns it, or empty if none was found.

### `property.Get` Mandatory Params

- `property.FilterParams.ID`: The string to be used as ID to get the desired property

### `property.Get` Optional Params & Default Values

This func has no optional params.

## `GetBySlug(params FilterParams) (common.Property, error)`

```go
propertyClient, err := property.NewClient(common.ClientParams{
  Credentials: common.Creds{ ... },
})

property, _ := propertyClient.GetBySlug(property.FilterParams{Slug: "cart-ca"})

fmt.Println(property) // { "id": "some-id", "slug": "some-slug", "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z" }
```

This func returns the first property in the list that matches the slug, or nil if no properties match the slug.

### `property.GetBySlug` Mandatory Params

- `property.FilterParams.Slug`: The string to be used as slug to filter the property list

<p align="right"><em><a href="../#edgioproperty">back to the main README</a></em></p>
