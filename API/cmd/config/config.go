package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

type Configuration struct {
	ConnectionString string `env:"DATA_SOURCE" default:"host=localhost user=postgres password=root dbname=hospitaldb port=5432 sslmode=disable TimeZone=Asia/bangkok"`
	SecretKey        string `env:"SECRET_KEY" default:"123-secret-key"`
}

func New() Configuration {

	if len(os.Args) == 2 {
		return readConfigFromArgument()
	}

	conf := Configuration{}
	v := reflect.ValueOf(&conf).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		envKey := fieldType.Tag.Get("env")
		envValue, ok := os.LookupEnv(envKey)

		switch ok {
		case true:
			field.SetString(envValue)
		case false:
			field.SetString(fieldType.Tag.Get("default"))
		}
	}

	return conf
}

func readConfigFromArgument() Configuration {
	conf := Configuration{}
	configPath, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalf("Error getting absolute path for config file: %s", err)
	}

	configBytes, err := ioutil.ReadFile(configPath)

	err = json.Unmarshal(configBytes, &conf)

	if err != nil {
		log.Fatalf("Error Unmarshal config file: %s", err)
	}

	return conf
}
