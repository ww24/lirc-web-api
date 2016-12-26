package config

import "testing"

func TestIsProd(t *testing.T) {
	// default is not production
	if IsProd() != false {
		t.Fatalf("IsProd(): expected: %v, actual: %v", false, IsProd())
	}
}

func TestIsDev(t *testing.T) {
	// default is not production
	if IsDev() != true {
		t.Fatalf("IsDev(): expected: %v, actual: %v", true, IsDev())
	}
}
