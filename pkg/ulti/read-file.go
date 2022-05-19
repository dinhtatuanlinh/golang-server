package ulti

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadFile(path string) (result *interface{}, err error) {
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
