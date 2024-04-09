package variable_test

import (
	"edgio/common"
	"edgio/variable"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetByKey(t *testing.T) {
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

	t.Run("should return the variable identified by the provided key", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(envVarsURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(variablesResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL}

		client, _ := variable.NewClient(params)
		variable, err := client.GetByKey(variable.FilterParams{Key: "some-env-var-key", EnvID: "1"})

		require.NoError(t, err)
		assert.Equal(t, "some-env-var-key", variable.Key)
	})

	t.Run("should error out if EnvID is missing", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(envVarsURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(variablesResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL}
		client, _ := variable.NewClient(params)

		_, err := client.GetByKey(variable.FilterParams{Key: "some-env-name"})

		assert.Equal(t, errors.New("'EnvID' is required"), err)
	})

	t.Run("should error out if Key is missing", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(envVarsURL, func(rw http.ResponseWriter, _ *http.Request) {
			_, err := rw.Write([]byte(variablesResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

		params.Config = common.ClientConfig{URL: server.URL}
		client, _ := variable.NewClient(params)

		_, err := client.GetByKey(variable.FilterParams{EnvID: "1"})

		assert.Equal(t, errors.New("'Key' is required"), err)
	})

	t.Run("should return an error if the list request fails", func(t *testing.T) {
		mux := http.NewServeMux()

		server := httptest.NewServer(mux)
		defer server.Close()

		mux.HandleFunc(envVarsURL, func(rw http.ResponseWriter, _ *http.Request) {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		})

		params.Config = common.ClientConfig{URL: server.URL}

		client, _ := variable.NewClient(params)
		_, err := client.GetByKey(variable.FilterParams{Key: "some-env-var-key", EnvID: "1"})

		require.Error(t, err)
		assert.Equal(t, "[HTTP ERROR]: Status Code: 500 - Internal Server Error", err.Error())
	})
}
