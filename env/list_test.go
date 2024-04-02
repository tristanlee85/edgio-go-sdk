package env_test

import (
	"edgio/common"
	"edgio/env"
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

	mux.HandleFunc(envsURL, func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(environmentsResponse))
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
		Config: common.ClientConfig{URL: server.URL},
	}

	client, _ := env.NewClient(params)
	result, _ := client.List("some-property-id")

	assert.Len(t, result.Items, 2)
	assert.Equal(t, "some-env-name", result.Items[0].Name)
	assert.Equal(t, "another-env-name", result.Items[1].Name)
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
		Config: common.ClientConfig{URL: ":"},
	}

	client, _ := env.NewClient(params)

	_, err := client.List("some-property-id")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "parse \":/accounts/v0.1/environments\": missing protocol scheme")
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
		Config: common.ClientConfig{URL: server.URL},
	}

	client, _ := env.NewClient(params)
	_, err := client.List("\n")

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

	mux.HandleFunc(envsURL, func(rw http.ResponseWriter, _ *http.Request) {
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
		Config: common.ClientConfig{URL: server.URL},
	}

	client, _ := env.NewClient(params)
	_, err := client.List("some-property-id")

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
		Config: common.ClientConfig{URL: server.URL},
	}

	mux.HandleFunc(envsURL, func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"items": "invalid"}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	client, _ := env.NewClient(params)
	_, err := client.List("some-property-id")

	require.Error(t, err)
	assert.Equal(t, "1 error(s) decoding:\n\n* 'items': source data must be an array or slice, got string", err.Error())
}

func TestFilterList(t *testing.T) {
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

	mux.HandleFunc(envsURL, func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(environmentsResponse))
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
		Config: common.ClientConfig{URL: server.URL},
	}

	client, _ := env.NewClient(params)
	filterParams := env.FilterParams{
		PropertyID: "some-property-id",
		Name:       "some-env-name",
	}
	result, _ := client.FilterList(filterParams)

	assert.Len(t, result.Items, 1)
	assert.Equal(t, "some-env-name", result.Items[0].Name)
}

func TestFilterListNoPropertyID(t *testing.T) {
	params := common.ClientParams{
		Credentials: common.Creds{
			Key:    "key",
			Secret: "secret",
			Scopes: "scopes",
		},
		Config: common.ClientConfig{URL: "http://localhost"},
	}

	client, _ := env.NewClient(params)
	filterParams := env.FilterParams{
		PropertyID: "",
		Name:       "some-env-name",
	}
	_, err := client.FilterList(filterParams)

	require.Error(t, err)
	assert.Equal(t, "PropertyID is required", err.Error())
}

func TestFilterListNoName(t *testing.T) {
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

	mux.HandleFunc(envsURL, func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(environmentsResponse))
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
		Config: common.ClientConfig{URL: server.URL},
	}

	client, _ := env.NewClient(params)
	filterParams := env.FilterParams{
		PropertyID: "some-property-id",
	}
	result, _ := client.FilterList(filterParams)

	assert.Len(t, result.Items, 2)
	assert.Equal(t, "some-env-name", result.Items[0].Name)
	assert.Equal(t, "another-env-name", result.Items[1].Name)
}

func TestFilterListListError(t *testing.T) {
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

	mux.HandleFunc(envsURL, func(rw http.ResponseWriter, _ *http.Request) {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	})

	params := common.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server2.URL,
		},
		Config: common.ClientConfig{URL: server.URL},
	}

	client, _ := env.NewClient(params)
	filterParams := env.FilterParams{
		PropertyID: "some-property-id",
		Name:       "some-env-name",
	}
	_, err := client.FilterList(filterParams)

	require.Error(t, err)
	assert.Equal(t, "[HTTP ERROR]: Status Code: 500 - Internal Server Error", err.Error())
}
