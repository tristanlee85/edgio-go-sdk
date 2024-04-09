package variable_test

const envVarsURL = "/config/v0.1/environment-variables"
const envVarURL = "/config/v0.1/environment-variables/some-id"
const authResult = `{"access_token": "test_token"}`
const variableResponse = `{
    "id": "some-id",
    "key": "some-env-var-key",
    "value": "some-value",
    "secret": true,
    "create_at": "2019-08-24T14:15:22Z",
    "updated_at": "2019-08-24T14:15:22Z"
}`
const variablesResponse = `{
    "total_items": 2,
    "items": [
        {
            "id": "some-id",
            "key": "some-env-var-key",
            "value": "some-value",
            "secret": true,
            "create_at": "2019-08-24T14:15:22Z",
            "updated_at": "2019-08-24T14:15:22Z"
        },
        {
            "id": "another-id",
            "key": "another-env-var-key",
            "value": "another-value",
            "secret": false,
            "create_at": "2019-08-24T14:15:22Z",
            "updated_at": "2019-08-24T14:15:22Z"
        }
    ]
}`
