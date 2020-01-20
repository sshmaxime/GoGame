package network

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ReadBody(rBody io.Reader, output interface{}) error {
	body, err := ioutil.ReadAll(rBody)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	return nil
}
