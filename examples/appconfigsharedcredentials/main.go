package main

import (
	"fmt"
	configs "github.com/sinhashubham95/go-config-client"
	"os"
)

func main() {
	client, _ := configs.New(configs.Options{
		Provider: configs.AWSAppConfig,
		Params: map[string]interface{}{
			"id":              "example-go-config-client",
			"region":          os.Getenv("region"),
			"app":             "sample",
			"env":             "sample",
			"configType":      "json",
			"configNames":     []string{"sample"},
			"credentialsMode": configs.AppConfigSharedCredentialMode,
		},
	})
	defer func() {
		_ = client.Close()
	}()

	fmt.Println(client.GetString("sample", "k1"))
	fmt.Println(client.GetString("sample", "k2"))
}
