package main

import (
	"testing"

	"github.com/nats-io/nats"
)

func TestConfigs(t *testing.T) {
	// Get configs
	_, err := getConfig()
	if err != nil {
		t.Error("Configs import error")
	}
}

func TestNats(t *testing.T) {
	conf, _ := getConfig()
	nc, err := nats.Connect(conf.NatsURL)
	if err != nil {
		t.Error("Nats connect err")
	}
	defer nc.Close()
}
