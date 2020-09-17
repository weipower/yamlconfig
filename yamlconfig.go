package yamlconfig

import (
    yaml "gopkg.in/yaml.v2"
    "io/ioutil"
)

func GetYamlFile(yamlfile string, yamlInf interface{}) error {
    yamlFile, err := ioutil.ReadFile(yamlfile)
    if err != nil {
        return err
    }
    err = yaml.Unmarshal(yamlFile, yamlInf)
    if err != nil {
        return err
    }
    return nil
}
func AllFieldsAreSet(yamlInf interface{}) (err error) {
    return CheckAllFieldsAreSet(yamlInf)
}
