package lirc

import "testing"

func TestList(t *testing.T) {
	client, err := NewClient("test/lircd")
	if err != nil {
		t.Fatal(err)
	}
	client.Verbose = true
	defer client.Close()

	_, err = client.List("")
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.List("lighting")
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.List("lighting", "up", "down")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendOnce(t *testing.T) {
	client, err := NewClient("test/lircd")
	if err != nil {
		t.Fatal(err)
	}
	client.Verbose = true
	defer client.Close()

	err = client.SendOnce("lighting", "on_off")
	if err != nil {
		t.Fatal(err)
	}
}
