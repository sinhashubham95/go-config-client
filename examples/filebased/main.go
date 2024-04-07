package main

import (
	"fmt"
	configs "github.com/sinhashubham95/go-config-client"
)

type decoded struct {
	A string `json:"a"`
	B string `json:"b"`
}

func main() {
	client, _ := configs.New(configs.Options{
		Provider: configs.FileBased,
		Params: map[string]interface{}{
			"configsDirectory": "./examples/filebased",
			"configNames":      []string{"configs"},
			"configType":       "json",
		},
	})
	defer func() {
		_ = client.Close()
	}()

	fmt.Println(client.GetD("configs", "a", "naruto"))
	fmt.Println(client.GetIntD("configs", "c", 5678))
	fmt.Println(client.GetFloatD("configs", "d", 5678.9))
	fmt.Println(client.GetStringD("configs", "a", "naruto"))
	fmt.Println(client.GetBoolD("configs", "e", false))

	d := decoded{}
	fmt.Println(client.Unmarshal("configs", "f", &d), d)
}
