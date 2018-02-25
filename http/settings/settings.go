package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var environments = map[string]string{
	"production":  "settings/prod.json",
	"development": "settings/dev.json",
	"test":        "../../../settings/test.json",
}

var settings interface{}
var env = "development"

func Init() {
	env = os.Getenv("IMAGEPROXY_ENV")

	if env == "" {
		fmt.Println("Warning: Setting development environment due to lack of IMAGEPROXY_ENV value")
		env = "development"
	}

	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])

	if err != nil {
		fmt.Println("Error while reading config file", err)
	}

	jsonErr := json.Unmarshal(content, &settings)

	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() map[string]interface{} {
	if &settings == nil {
		Init()
	}
	return settings.(map[string]interface{})
}

func IsProductionEnvironment() bool {
	return env == "production"
}

func IsDevelopmentEnvironment() bool {
	return env == "development"
}

func IsTestEnvironment() bool {
	return env == "test"
}
