package main

import "testing"

func TestFetchSignals(t *testing.T) {
	signals, err := fetchSignals("")
	if err != nil {
		t.Fatal(err)
	}

	if len(signals) != 5 {
		t.Fatalf("len(signals): expected: %d, actual: %d", 5, len(signals))
	}

	expectedSignals := []string{"up", "down", "on_off", "pilot", "memory"}
	for i, signal := range signals {
		if signal.Remote != "lighting" {
			t.Fatalf("signal.Remote: expected: %s, actual: %s", "lighting", signal.Remote)
		}
		if signal.Name != expectedSignals[i] {
			t.Fatalf("signal.Name: expected: %s, actual: %s", expectedSignals[i], signal.Name)
		}
	}
}

func TestSendSignal(t *testing.T) {
	sig := &signal{
		Remote: "lighting",
		Name:   "up",
	}
	err := sendSignal(sig)
	if err != nil {
		t.Fatal(err)
	}
}
