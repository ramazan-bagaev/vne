package main

import (
	"testing"
)

func TestParseCommandCreate(t *testing.T) {
	cmd := ParseCommand([]string{"vne", "create", "-u", "user", "-d", "/some/place"})

	if cmd.Cmd != "create" || cmd.ConfigPath != "/some/place" || cmd.User != "user" {
		t.Errorf("parse failed %s", cmd)
	}
}

func TestParseCommandCreateDefaults(t *testing.T) {
	cmd := ParseCommand([]string{"vne", "create"})

	if cmd.Cmd != "create" || cmd.ConfigPath != "" || cmd.User != "vne-user" {
		t.Errorf("parse failed %s", cmd)
	}
}
