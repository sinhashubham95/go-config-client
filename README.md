# Go Config Client

This is the config client for Go projects.

- **File-Based Client** - Fully Implemented
- **ETCD Client** - TBD
- **AWS App Config Client** - Fully Implemented

## Project Versioning

Go Config Client uses [semantic versioning](http://semver.org/). API should not change between patch and minor releases. New minor versions may add additional features to the API.

## Installation

To install `Go Config Client` package, you need to install Go and set your Go workspace first.

1. The first need Go installed (version 1.13+ is required), then you can use the below Go command to install Go Config Client.

```shell
go get github.com/sinhashubham95/go-config-client
```

2. Because this is a private repository, you will need to mark this in the Go env variables.

```shell
go env -w GOPRIVATE=github.com/sinhashubham95/*
```

3. Also, follow this to generate a personal access token and add the following line to your $HOME/.netrc file.

```
machine github.com login ${USERNAME} password ${PERSONAL_ACCESS_TOKEN}
```

4. Import it in your code:

```go
import configs "github.com/sinhashubham95/go-config-client"
```

## Usage

### New Client

```go
import configs "github.com/sinhashubham95/go-config-client"

fileBasedClient, _ := configs.New(configs.Options{
    Provider: configs.FileBased,
    Params: map[string]interface{}{
        "configsDirectory": ".",
        "configNames": []string{"configs"},
        "configType": "json",
    },
})

awsAppConfigClient, _ := configs.New(configs.Options{
	Provider: configs.AWSAppConfig,
	Params: map[string]interface{}{
        "id":          "example-go-config-client",
        "region":      os.Getenv("region"),
        "accessKeyId": os.Getenv("accessKeyId"),
        "secretKey":   os.Getenv("secretKey"),
        "app":         "sample",
        "env":         "sample",
        "configType":  "json",
        "configNames": []string{"sample"},
    }
})
```

### Configuring Secrets Manager

1. Add additional param secretNames as shown below
2. App Config and SecretsManager will use same access and secret keys hence makes sure necessary access is available 
3. Secrets values are fetched on startup time and cached by library. Restart needed if changing secrets.
4. Secret data in case of AWS should contain JSON data only.

```go
import configs "github.com/sinhashubham95/go-config-client"

fileBasedClient, _ := configs.New(configs.Options{
    Provider: configs.FileBased,
    Params: map[string]interface{}{
        "configsDirectory": ".",
        "configNames": []string{"configs"},
        "configType": "json",
        "secretsDirectory": ".",
        "secretNames": []string{"secrets"},
        "secretType": "json",
    },
})

awsAppConfigClient, _ := configs.New(configs.Options{
	Provider: configs.AWSAppConfig,
	Params: map[string]interface{}{
        "id":          "example-go-config-client",
        "region":      os.Getenv("region"),
        "accessKeyId": os.Getenv("accessKeyId"),
        "secretKey":   os.Getenv("secretKey"),
        "app":         "sample",
        "env":         "sample",
        "configType":  "json",
        "configNames": []string{"sample"},
        "secretNames": []string{"SecretName"},
    }
})
```
### Getting Secrets Value

1. New methods have been exposed to fetch secrets GetSecret, GetIntSecret, GetFloatSecret and GetStringSecret.
2. Secrets values are fetched on startup time and cached by library. Restart needed if changing secrets.

### Getting Configs

There are 2 types of methods available.

1. Plain methods which take the config name and the key.
2. Methods with default values which take the config name, key and the default value. The default value will be used in case the value is not found in the config mentioned corresponding to the key asked for.

### Note

For **File Based Config Client**, the config name is the name of the file from where the configurations have to be referenced, and the key is the location of the config being fetched from that configuration file.

For **AWS App Config Client**, the config name is the name of the configuration profile deployed, and the key is the location of the config being fetched from that configuration profile.
