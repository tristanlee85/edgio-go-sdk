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

	t.Run("should return the environment identified by provided ID", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(envURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(environmentResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL}

		client, _ := env.NewClient(params)
		env, _ := client.Get(env.FilterParams{ID: "some-id"})

		assert.Equal(t, "some-env-name", env.Name)
	})

	t.Run("should return empty if provided ID does not match any env", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(envURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(environmentResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL}

		client, _ := env.NewClient(params)
		env, _ := client.Get(env.FilterParams{ID: "unmatched-id"})

		assert.Empty(t, env.Name)
	})

	t.Run("should return error if the URL is not parseable", func(t *testing.T) {
		params.Config = common.ClientConfig{URL: ":"}

		client, _ := env.NewClient(params)
		_, err := client.Get(env.FilterParams{ID: "some-id"})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "parse \":/accounts/v0.1/environments/some-id\": missing protocol scheme")
	})

	t.Run("should return error if ID is not provided", func(t *testing.T) {
		client, _ := env.NewClient(params)
		_, err := client.Get(env.FilterParams{})

		require.Error(t, err)
		assert.Equal(t, "'ID' is required", err.Error())
	})

	t.Run("should return error if mapstructure decode fails", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(envURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(`{"ID": true}`))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL}

		client, _ := env.NewClient(params)
		_, err := client.Get(env.FilterParams{ID: "some-id"})

		require.Error(t, err)
		assert.Equal(
			t,
			"1 error(s) decoding:\n\n* 'id' expected type 'string', got unconvertible type 'bool', value: 'true'",
			err.Error(),
		)
	})
}
