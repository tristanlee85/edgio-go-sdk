package property_test

const propertiesURL = "/accounts/v0.1/properties"
const propertyURL = "/accounts/v0.1/properties/some-id"
const authResult = `{"access_token": "test_token"}`
const propertyResponse = `{
    "id": "some-id",
    "slug": "some-slug",
    "created_at": "2019-08-24T14:15:22Z",
    "updated_at": "2019-08-24T14:15:22Z"
}`
const propertiesResponse = `{
    "total_items": 2,
    "items": [
        {
            "id": "some-id",
            "slug": "some-slug",
            "created_at": "2019-08-24T14:15:22Z",
            "updated_at": "2019-08-24T14:15:22Z"
        },
        {
            "id": "another-id",
            "slug": "another-slug",
            "created_at": "2019-08-24T14:15:22Z",
            "updated_at": "2019-08-24T14:15:22Z"
        }
    ]
}`
