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

var credentials = common.Creds{
	Key:    os.Getenv("EDGIO_API_CLIENT_ID"),
	Secret: os.Getenv("EDGIO_API_CLIENT_SECRET"),
}
var orgClient, _ = org.NewClient(common.ClientParams{Credentials: credentials})
var organization, _ = orgClient.Get(common.URLParams{Path: os.Getenv("EDGIO_ORG_ID")})
var clientParams = common.ClientParams{
	Credentials: credentials,
	Config:      common.ClientConfig{OrgID: organization.ID, AccessToken: orgClient.Client.AccessToken},
}
var propertyClient, _ = property.NewClient(clientParams)
var envClient, _ = env.NewClient(clientParams)
var variableClient, _ = variable.NewClient(clientParams)

func propertyFilterListCkech() []common.Property {
	fmt.Println("FilterList (Property)")
	properties, _ := propertyClient.FilterList(property.FilterParams{Slug: "another-"})

	for _, propperty := range properties.Items {
		fmt.Println("Property from FilterList: " + propperty.Slug)
	}

	return properties.Items
}

func envFilterListCheck(PropertyID string) []common.Env {
	fmt.Println("FilterList (Env)")
	envs, _ := envClient.FilterList(env.FilterParams{PropertyID: PropertyID, Name: "another"})

	for _, env := range envs.Items {
		fmt.Println("Env from FilterList: " + env.Name)
	}

	return envs.Items
}

func variableFilterListCheck(EnvID string) []common.Variable {
	fmt.Println("FilterList (Variable)")
	variables, _ := variableClient.FilterList(variable.FilterParams{EnvID: EnvID, Key: "SOME"})

	for _, variable := range variables.Items {
		fmt.Println("Variable from FilterList: " + variable.Key)
	}

	return variables.Items
}

func propertyGetCheck(ID string) common.Property {
	fmt.Println("Property Get (by id)")
	fmt.Println("Property ID: " + ID)
	propGetResult, _ := propertyClient.Get(property.FilterParams{ID: ID})
	fmt.Println("Property Get Result: " + propGetResult.Slug)

	return propGetResult
}

func getByAttrCheck() {
	fmt.Println("GetByAttr")

	property, _ := propertyClient.GetBySlug(property.FilterParams{Slug: "some-property"})

	fmt.Println("Property GetBySlug Result: " + property.Slug)

	stageEnv, _ := envClient.GetByName(env.FilterParams{PropertyID: property.ID, Name: "some-env"})

	fmt.Println("Env GetByName Result: " + stageEnv.Name)

	variable, _ := variableClient.GetByKey(variable.FilterParams{EnvID: stageEnv.ID, Key: "SOME_ENV_VAR"})

	fmt.Println("Variable GetByKey Result: " + variable.Key + " = " + variable.Value)

}

func main() {
	fmt.Println("main.go")

	fmt.Println("Org ID: " + organization.ID)

	properties := propertyFilterListCkech()

	for _, prop := range properties {
		envs := envFilterListCheck(prop.ID)

		for _, env := range envs {
			variableFilterListCheck(env.ID)
		}

		propertyGetCheck(prop.ID)
	}

	getByAttrCheck()

	fmt.Println("main.go")
}
