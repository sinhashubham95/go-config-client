package main

import (
	"fmt"
	configs "github.com/sinhashubham95/go-config-client"
	"os"
	"time"
)

type decoded struct {
	A int      `yaml:"a"`
	B int      `yaml:"b"`
	C int      `yaml:"c"`
	D []string `yaml:"d"`
}

func main() {
	client, _ := configs.New(configs.Options{
		Provider: configs.AWSAppConfig,
		Params: map[string]interface{}{
			"id":          "example-go-config-client",
			"region":      os.Getenv("region"),
			"accessKeyId": os.Getenv("accessKeyId"),
			"secretKey":   os.Getenv("secretKey"),
			"app":         "aws-config-deploy-example",
			"env":         "dev",
			"configType":  "yaml",
			"configNames": []string{"b"},
		},
	})
	defer func() {
		_ = client.Close()
	}()

	fmt.Println(client.GetString("b", "naruto"))
	fmt.Println(client.GetString("b", "boruto"))

	d := decoded{}
	fmt.Println(client.Unmarshal("b", "nested", &d), d)

	_ = client.AddChangeListener("a", func(params ...interface{}) {
		fmt.Println("configuration changed", params[0])
		fmt.Println(client.GetString("a", "naruto"))
		fmt.Println(client.GetString("a", "boruto"))
	})

	time.Sleep(time.Minute * 5)
}
