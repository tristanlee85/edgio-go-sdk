package env_test

import (
	"edgio/common"
	"edgio/env"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetByName(t *testing.T) {
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

	t.Run("should return the env identified by the provided name", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(envsURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(environmentsResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL}

		client, _ := env.NewClient(params)
		env, err := client.GetByName(env.FilterParams{PropertyID: "some-property-id", Name: "some-env-name"})

		require.NoError(t, err)
		assert.Equal(t, "some-env-name", env.Name)
	})

	t.Run("should error out if PropertyID is missing", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(envsURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(environmentsResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL}
		client, _ := env.NewClient(params)

		_, err := client.GetByName(env.FilterParams{Name: "some-env-name"})

		assert.Equal(t, errors.New("'PropertyID' is required"), err)
	})

	t.Run("should error out if Name is missing", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(envsURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(environmentsResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL}
		client, _ := env.NewClient(params)

		_, err := client.GetByName(env.FilterParams{PropertyID: "1"})

		assert.Equal(t, errors.New("'Name' is required"), err)
	})

	t.Run("should return an error if the list request fails", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(envsURL, func(rw http.ResponseWriter, _ *http.Request) {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		})

		params.Config = common.ClientConfig{URL: server.URL}

		client, _ := env.NewClient(params)
		_, err := client.GetByName(env.FilterParams{PropertyID: "some-property-id", Name: "some-env-name"})

		require.Error(t, err)
		assert.Equal(t, "[HTTP ERROR]: Status Code: 500 - Internal Server Error", err.Error())
	})
}
