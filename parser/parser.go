package parser

import (
	"encoding/json"
	"io/ioutil"
)

func Parse(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	res := &Config{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
