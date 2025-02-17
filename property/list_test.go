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

func TestList(t *testing.T) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)
	defer server.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(authResult))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server2.Close()

	mux.HandleFunc(propertiesURL, func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(propertiesResponse))
		if err != nil {
			t.Fatal(err)
		}
	})

	params := common.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server2.URL,
		},
		Config: common.ClientConfig{URL: server.URL, OrgID: "some-org-id"},
	}

	client, _ := property.NewClient(params)
	result, _ := client.List()

	assert.Len(t, result.Items, 2)
	assert.Equal(t, "some-slug", result.Items[0].Slug)
	assert.Equal(t, "another-slug", result.Items[1].Slug)
}

func TestListParseURLError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(authResult))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	params := common.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server.URL,
		},
		Config: common.ClientConfig{URL: ":", OrgID: "some-org-id"},
	}

	client, _ := property.NewClient(params)

	_, err := client.List()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "parse \":/accounts/v0.1/properties\": missing protocol scheme")
}

func TestListNewRequestError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(authResult))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	params := common.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server.URL,
		},
		Config: common.ClientConfig{URL: server.URL, OrgID: "\n"},
	}

	client, _ := property.NewClient(params)
	_, err := client.List()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid control character in URL")
}

func TestListGetHTTPJSONResultError(t *testing.T) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)
	defer server.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(authResult))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server2.Close()

	mux.HandleFunc(propertiesURL, func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`error`))
		if err != nil {
			t.Fatal(err)
		}
	})

	params := common.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server2.URL,
		},
		Config: common.ClientConfig{URL: server.URL, OrgID: "some-org-id"},
	}

	client, _ := property.NewClient(params)
	_, err := client.List()

	require.Error(t, err)
	assert.Equal(t, "invalid character 'e' looking for beginning of value", err.Error())
}

func TestListMapstructureDecodeError(t *testing.T) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)
	defer server.Close()

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
		Config: common.ClientConfig{
			URL:   server.URL,
			OrgID: "some-org",
		},
	}

	mux.HandleFunc(propertiesURL, func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"items": "invalid"}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	client, _ := property.NewClient(params)
	_, err := client.List()

	require.Error(t, err)
	assert.Equal(t, "1 error(s) decoding:\n\n* 'items': source data must be an array or slice, got string", err.Error())
}

func TestFilterList(t *testing.T) {
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
		Config: common.ClientConfig{OrgID: "some-org-id"},
	}

	t.Run("returns error when List fails", func(t *testing.T) {
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		params.Config.URL = server.URL

		defer server.Close()

		mux.HandleFunc(propertiesURL, func(rw http.ResponseWriter, _ *http.Request) {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		})

		client, _ := property.NewClient(params)
		_, err := client.FilterList(property.FilterParams{})

		assert.Equal(t, "[HTTP ERROR]: Status Code: 500 - Internal Server Error", err.Error())
		require.Error(t, err)
	})

	t.Run("returns full list when no slug is provided", func(t *testing.T) {
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		params.Config.URL = server.URL

		defer server.Close()

		mux.HandleFunc(propertiesURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(propertiesResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		client, _ := property.NewClient(params)
		result, err := client.FilterList(property.FilterParams{})

		require.NoError(t, err)
		assert.Equal(t, 2, result.Total)
		assert.Equal(t, 2, result.FilteredTotal)
	})

	t.Run("returns filtered list when slug is provided", func(t *testing.T) {
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		params.Config.URL = server.URL

		defer server.Close()

		mux.HandleFunc(propertiesURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(propertiesResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		client, _ := property.NewClient(params)
		result, err := client.FilterList(property.FilterParams{Slug: "another-slug"})

		require.NoError(t, err)
		assert.Equal(t, 2, result.Total)
		assert.Equal(t, 1, result.FilteredTotal)
		assert.Equal(t, "another-slug", result.Items[0].Slug)
	})
}
