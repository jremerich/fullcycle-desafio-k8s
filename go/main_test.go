package main

import "testing"

func TestGreeting(t *testing.T) {
	result := greeting("Test")
	if result != "<b>Test</b>" {
		t.Errorf("greeting failed: expected <b>Test</b> | got %s", result)
	}
}