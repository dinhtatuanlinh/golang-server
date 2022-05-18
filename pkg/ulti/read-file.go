package ulti

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadFile(path string) (obj *interface{}, err error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(buf, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}