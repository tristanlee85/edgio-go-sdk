package env_test

const envsURL = "/accounts/v0.1/environments"
const authResult = `{"access_token": "test_token"}`
const environmentsResponse = `{
    "total_items": 2,
    "items": [
        {
            "id": "some-id",
            "name": "some-env-name",
            "legacy_account_number": "",
            "default_domain_name": "",
            "dns_domain_name": "",
            "can_members_deploy": true,
            "only_maintainers_can_deploy": true,
            "http_request_logging": true,
            "pci_compliance": true,
            "created_at": "2019-08-24T14:15:22Z",
            "updated_at": "2019-08-24T14:15:22Z"
        },
        {
            "id": "another-id",
            "name": "another-env-name",
            "legacy_account_number": "",
            "default_domain_name": "",
            "dns_domain_name": "",
            "can_members_deploy": false,
            "only_maintainers_can_deploy": false,
            "http_request_logging": false,
            "pci_compliance": false,
            "created_at": "2019-08-24T14:15:22Z",
            "updated_at": "2019-08-24T14:15:22Z"
        }
    ]
}`
