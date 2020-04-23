package parser

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func load(path string) ([]byte, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func Parse(path string, val interface{}) error {
	content, err := load(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, val)
}
