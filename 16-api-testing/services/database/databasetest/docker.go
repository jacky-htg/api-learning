package databasetest

import (
	"bytes"
	"os/exec"
	"testing"
)

// StartContainer runs a mysql container to execute commands.
func StartContainer(t *testing.T) {
	t.Helper()

	cmd := exec.Command("docker-compose", "up", "-d")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		t.Fatalf("could not start docker-compose: %v", err)
	}
}

// StopContainer stops and removes the specified container.
func StopContainer(t *testing.T) {
	t.Helper()

	if err := exec.Command("docker-compose", "down").Run(); err != nil {
		t.Fatalf("could not stop docker-compose: %v", err)
	}
}
