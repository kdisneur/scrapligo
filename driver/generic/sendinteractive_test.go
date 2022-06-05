package generic_test

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"

	"github.com/scrapli/scrapligo/channel"

	"github.com/google/go-cmp/cmp"
)

type sendInteractiveTestCase struct {
	description      string
	events           []*channel.SendInteractiveEvent
	completePatterns []*regexp.Regexp
	payloadFile      string
}

func testSendInteractive(testName string, testCase *sendInteractiveTestCase) func(t *testing.T) {
	return func(t *testing.T) {
		t.Logf("%s: starting", testName)

		d, fileTransportObj := prepareDriver(t, testName, testCase.payloadFile)

		r, err := d.SendInteractive(testCase.events)
		if err != nil {
			t.Errorf(
				"%s: encountered error running generic Driver GetPrompt, error: %s",
				testName,
				err,
			)
		}

		if r.Failed != nil {
			t.Fatalf("%s: response object indicates failure",
				testName)
		}

		actualOut := r.Result
		actualIn := bytes.Join(fileTransportObj.Writes, []byte("\n"))

		if *update {
			writeGolden(t, testName, actualIn, actualOut)
		}

		expectedIn := readFile(t, fmt.Sprintf("golden/%s-in.txt", testName))
		expectedOut := readFile(t, fmt.Sprintf("golden/%s-out.txt", testName))

		if !cmp.Equal(actualIn, expectedIn) {
			t.Fatalf(
				"%s: actual and expected inputs do not match\nactual: %s\nexpected:%s",
				testName,
				actualIn,
				expectedIn,
			)
		}

		if !cmp.Equal(actualOut, string(expectedOut)) {
			t.Fatalf(
				"%s: actual and expected outputs do not match\nactual: %s\nexpected:%s",
				testName,
				actualOut,
				expectedOut,
			)
		}
	}
}

func TestSendInteractive(t *testing.T) {
	cases := map[string]*sendInteractiveTestCase{
		"send-interactive-simple": {
			description: "simple send interactive test",
			events: []*channel.SendInteractiveEvent{
				{
					ChannelInput:    "clear logging",
					ChannelResponse: "[confirm]",
					HideInput:       false,
				},
				{
					ChannelInput:    "",
					ChannelResponse: "",
					HideInput:       false,
				},
			},
			completePatterns: nil,
			payloadFile:      "send-interactive-simple.txt",
		},
	}

	for testName, testCase := range cases {
		f := testSendInteractive(testName, testCase)

		t.Run(testName, f)
	}
}
