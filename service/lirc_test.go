package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ww24/lirc-web-api/lirc"
)

func TestFetchSignals(t *testing.T) {
	{
		ctrl := gomock.NewController(t)
		mockLIRC := lirc.NewMockClientAPI(ctrl)
		newLIRCService = func(...string) (lirc.ClientAPI, error) {
			return mockLIRC, nil
		}

		gomock.InOrder(
			mockLIRC.EXPECT().List(
				"",
			).Return(
				[]string{"lighting"},
				nil,
			),
			mockLIRC.EXPECT().List(
				"lighting",
			).Return(
				[]string{
					"0000000000000001 up",
					"0000000000000002 down",
					"0000000000000003 on_off",
					"0000000000000004 pilot",
					"0000000000000005 memory",
				},
				nil,
			),
			mockLIRC.EXPECT().Close(),
		)
	}

	{
		signals, err := FetchSignals("")
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
}

func TestSendSignal(t *testing.T) {
	{
		ctrl := gomock.NewController(t)
		mockLIRC := lirc.NewMockClientAPI(ctrl)
		newLIRCService = func(...string) (lirc.ClientAPI, error) {
			return mockLIRC, nil
		}

		gomock.InOrder(
			mockLIRC.EXPECT().List(
				"lighting",
				"up",
			).Return(
				[]string{"0000000000000001 up"},
				nil,
			),
			mockLIRC.EXPECT().SendOnce(
				"lighting",
				"up",
			).Return(
				nil,
			),
			mockLIRC.EXPECT().Close(),
		)
	}

	{
		sig := &Signal{
			Remote: "lighting",
			Name:   "up",
		}
		err := SendSignal(&SendSignalParam{Signal: sig})
		if err != nil {
			t.Fatal(err)
		}
	}
}
