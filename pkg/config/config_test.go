package config

import "testing"

func TestNewConfig(t *testing.T) {
	clusterConfig, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Worker nodes: %v\n", clusterConfig)
}
