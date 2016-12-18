package lirc

import "testing"

var (
	buttons = []string{
		"0000000000000001 up",
		"0000000000000002 down",
		"0000000000000003 on_off",
		"0000000000000004 pilot",
		"0000000000000005 memory",
	}
)

func TestList(t *testing.T) {
	client, err := NewClient("test/lircd")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	reps, err := client.List("")
	if err != nil {
		t.Fatal(err)
	}
	if len(reps) != 1 {
		t.Fatal("len(reps): expected: 1, actual:", len(reps))
	}
	if reps[0] != "lighting" {
		t.Fatal("reps[0]: expected: lighting, actual:", reps[0])
	}

	reps, err = client.List("lighting")
	if err != nil {
		t.Fatal(err)
	}
	if len(reps) != 5 {
		t.Fatal("len(reps): expected: 5, actual:", len(reps))
	}
	for i, rep := range reps {
		if rep != buttons[i] {
			t.Fatalf("reps[%d]: expected: %s, actual: %s", i, buttons[i], rep)
		}
	}

	reps, err = client.List("lighting", "up", "memory")
	if err != nil {
		t.Fatal(err)
	}
	if len(reps) != 2 {
		t.Fatal("len(reps): expected: 2, actual:", len(reps))
	}
	if reps[0] != buttons[0] {
		t.Fatalf("reps[0]: expected: %s, actual: %s", buttons[0], reps[0])
	}
	if reps[1] != buttons[4] {
		t.Fatalf("reps[1]: expected: %s, actual: %s", buttons[4], reps[0])
	}
}

func TestSendOnce(t *testing.T) {
	client, err := NewClient("test/lircd")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	err = client.SendOnce("lighting", "on_off")
	if err.Error() != "hardware does not support sending" {
		t.Fatal("err: expected: hardware does not support sending, actual:", err.Error())
	}
}
