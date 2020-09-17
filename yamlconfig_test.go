package yamlconfig

import (
	"encoding/json"
	"testing"
	"time"
)

type ConfigYaml struct {
	HOST    string `yaml:"HOST"`
	PORT    int    `yaml:"PORT"`
	SETTING struct {
		Redis struct {
			Db       int    `yaml:"db"`
			Host     string `yaml:"host"`
			Name     string `yaml:"name"`
			Password string `yaml:"password"`
			Prefix   string `yaml:"prefix"`
		} `yaml:"redis"`
	} `yaml:"SETTING"`
	URL   string    `yaml:"URL"`
	START time.Time `yaml:"START"`
}

var config ConfigYaml

func TestCmp(t *testing.T) {
	GetYamlFile("yamlconfig_test.yaml", &config)

	s, _ := json.Marshal(config)
	t.Log(string(s))

	err := AllFieldsAreSet(config)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
