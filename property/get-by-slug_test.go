package property_test

import (
	"edgio/common"
	"edgio/property"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBySlug(t *testing.T) {
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

	t.Run("should return the property identified by the provided slug", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(propertiesURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(propertiesResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL, OrgID: "some-org-id"}

		client, _ := property.NewClient(params)
		property, err := client.GetBySlug(property.FilterParams{Slug: "some-slug"})

		require.NoError(t, err)
		assert.Equal(t, "some-slug", property.Slug)
	})

	t.Run("should error out if Slug is missing", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(propertiesURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(propertiesResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL, OrgID: "some-org-id"}
		client, _ := property.NewClient(params)

		_, err := client.GetBySlug(property.FilterParams{})

		assert.Equal(t, errors.New("'Slug' is required"), err)
	})

	t.Run("should return an error if the list request fails", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(propertiesURL, func(rw http.ResponseWriter, _ *http.Request) {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		})

		params.Config = common.ClientConfig{URL: server.URL, OrgID: "some-org-id"}

		client, _ := property.NewClient(params)
		_, err := client.GetBySlug(property.FilterParams{Slug: "some-slug"})

		require.Error(t, err)
		assert.Equal(t, "[HTTP ERROR]: Status Code: 500 - Internal Server Error", err.Error())
	})
}
