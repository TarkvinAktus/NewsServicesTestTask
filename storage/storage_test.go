package main

import (
	"testing"
)

func TestConfigs(t *testing.T) {
	// Get configs
	_, err := getConfig()
	if err != nil {
		t.Error("Configs import error")
	}
}

func TestRedis(t *testing.T) {
	// Get configs
	conf, _ := getConfig()

	data := getDataFromRedis(conf, "any_key")

	if data != "some news text" {
		t.Error("Configs import error")
	}
}
