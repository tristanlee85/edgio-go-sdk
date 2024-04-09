package main

import (
	"edgio/common"
	"edgio/env"
	"edgio/org"
	"edgio/property"
	"edgio/variable"
	"fmt"
	"os"
)

func main() {
	fmt.Println("main.go")
	credentials := common.Creds{
		Key:    os.Getenv("EDGIO_API_CLIENT_ID"),
		Secret: os.Getenv("EDGIO_API_CLIENT_SECRET"),
	}

	orgClient, _ := org.NewClient(common.ClientParams{Credentials: credentials})
	org, _ := orgClient.Get(common.URLParams{Path: os.Getenv("EDGIO_ORG_ID")})

	fmt.Println("Org ID: " + org.ID)

	propertyClient, _ := property.NewClient(common.ClientParams{
		Credentials: credentials,
		Config:      common.ClientConfig{OrgID: org.ID, AccessToken: orgClient.Client.AccessToken},
	})

	fmt.Println("FilterList")
	properties, _ := propertyClient.FilterList(property.FilterParams{Slug: "another-"})

	envClient, _ := env.NewClient(common.ClientParams{
		Credentials: credentials,
		Config:      common.ClientConfig{AccessToken: orgClient.Client.AccessToken},
	})

	variableClient, _ := variable.NewClient(common.ClientParams{
		Credentials: credentials,
		Config:      common.ClientConfig{AccessToken: orgClient.Client.AccessToken},
	})

	for _, prop := range properties.Items {
		fmt.Println("Property from FilterList: " + prop.Slug)

		envs, _ := envClient.FilterList(env.FilterParams{PropertyID: prop.ID, Name: "another"})

		for _, env := range envs.Items {
			fmt.Println("Env from FilterList: " + env.Name)

			variables, _ := variableClient.FilterList(variable.FilterParams{EnvID: env.ID, Key: "SOME"})

			for _, variable := range variables.Items {
				fmt.Println("Variable from FilterList: " + variable.Key)
			}
		}

		fmt.Println("Property Get (by id)")
		fmt.Println("Property ID: " + prop.ID)
		propGetResult, _ := propertyClient.Get(property.FilterParams{ID: prop.ID})
		fmt.Println("Property Get Result: " + propGetResult.Slug)
	}

	fmt.Println("GetByAttr")

	property, _ := propertyClient.GetBySlug(property.FilterParams{Slug: "some-property"})

	fmt.Println("Property GetBySlug Result: " + property.Slug)

	stageEnv, _ := envClient.GetByName(env.FilterParams{PropertyID: property.ID, Name: "some-env"})

	fmt.Println("Env GetByName Result: " + stageEnv.Name)

	variable, _ := variableClient.GetByKey(variable.FilterParams{EnvID: stageEnv.ID, Key: "SOME_ENV_VAR"})

	fmt.Println("Variable GetByKey Result: " + variable.Key + " = " + variable.Value)

	fmt.Println("main.go")
}
