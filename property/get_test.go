package property_test

import (
	"edgio/common"
	"edgio/property"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(authResult))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server2.Close()

	params := common.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server2.URL,
		},
	}

	t.Run("should return the property identified by provided ID", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(propertyURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(propertyResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL, OrgID: "some-org-id"}

		client, _ := property.NewClient(params)
		property, _ := client.Get(property.FilterParams{ID: "some-id"})

		assert.Equal(t, "some-slug", property.Slug)
	})

	t.Run("should return empty if provided ID does not match any property", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(propertyURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(propertyResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL, OrgID: "some-org-id"}

		client, _ := property.NewClient(params)
		property, _ := client.Get(property.FilterParams{ID: "unmatched-id"})

		assert.Empty(t, property.Slug)
	})

	t.Run("should return error if the URL is not parseable", func(t *testing.T) {
		params.Config = common.ClientConfig{URL: ":", OrgID: "some-org-id"}

		client, _ := property.NewClient(params)
		_, err := client.Get(property.FilterParams{ID: "some-id"})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "parse \":/accounts/v0.1/properties/some-id\": missing protocol scheme")
	})

	t.Run("should return error if ID is not provided", func(t *testing.T) {
		client, _ := property.NewClient(params)
		_, err := client.Get(property.FilterParams{})

		require.Error(t, err)
		assert.Equal(t, "'ID' is required", err.Error())
	})

	t.Run("should return error if mapstructure decode fails", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(propertyURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(`{"ID": true}`))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL, OrgID: "some-org-id"}

		client, _ := property.NewClient(params)
		_, err := client.Get(property.FilterParams{ID: "some-id"})

		require.Error(t, err)
		assert.Equal(
			t,
			"1 error(s) decoding:\n\n* 'id' expected type 'string', got unconvertible type 'bool', value: 'true'",
			err.Error(),
		)
	})
}
