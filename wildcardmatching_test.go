package wildcardmatching

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"
)

type Test struct {
	Input string `json:"input"`
	Pattern string `json:"pattern"`
	Output bool  `json:"output"`
}

func TestWildCardMatching(t *testing.T) {
	f, err := os.Open("tests.json")

	if err != nil {
		t.Error(err)
		return
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			t.Error()
		}
	}(f)

	decoder := json.NewDecoder(bufio.NewReader(f))

	for {
		var tests map[string]Test

		err = decoder.Decode(&tests)

		if err == nil {
			for testName, test := range tests {
				t.Run(testName, func(tN *testing.T) {
					res := isMatch(test.Input, test.Pattern)

					if res != test.Output {
						tN.Fail()
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			t.Error(err)
			break
		}

	}
}
