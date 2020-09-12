package main

import (
	"fmt"
)

func InvalidRequest(msg string) error {
	return fmt.Errorf("invalid request:%v", msg)
}
