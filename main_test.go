package main

import (
	"gotest/system"
	"testing"
)

func TestClient(t *testing.T) {
	system.Run()
}

func TestServer(t *testing.T) {
	system.StartServer()
}
func TestForward(t *testing.T) {
	system.RunForward()
}
