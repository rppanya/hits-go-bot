package main

import (
	"testing"
)

func TestBotState(t *testing.T) {
	state := NewBotState()
	chatID := int64(12345)

	state.SetState(chatID, "awaiting_input")
	if state.GetState(chatID) != "awaiting_input" {
		t.Errorf("Expected 'awaiting_input', got %s", state.GetState(chatID))
	}

	state.SetState(chatID, "input_received")
	if state.GetState(chatID) != "input_received" {
		t.Errorf("Expected 'input_received', got %s", state.GetState(chatID))
	}
}
