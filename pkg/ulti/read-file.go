package ulti

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadFileYaml(path string) (result *interface{}, err error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		result = nil
		return
	}
	var obj interface{}
	err = yaml.Unmarshal(buf, &obj)
	
	result = &obj
	return
}

func ReadFileJson(path string) (result *interface{}, err error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		result = nil
		return
	}
	var obj interface{}
	err = json.Unmarshal(buf, &obj)
	
	result = &obj
	return
}