package main

import (
	"errors"
)

func InvalidRequest() error {
	return errors.New("invalid request")
}
