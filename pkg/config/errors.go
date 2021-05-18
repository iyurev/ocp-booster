package config

import "errors"

var (
	ErrNodeNotFound = errors.New("cluster node with such name isn't found")
)
