package parser

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type JsonParser struct{}

func NewJsonParser() *JsonParser {
	return &JsonParser{}
}

func (parser *JsonParser) GetType() string {
	return "json"
}

func (parser *JsonParser) Parse(r *http.Request, data interface{}) error {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, data); err != nil {
		return err
	}

	return nil
}
