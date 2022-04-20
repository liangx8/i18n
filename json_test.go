package i18n_test

import (
	"strings"
	"testing"

	"github.com/liangx8/i18n"
)

func Test_json(t *testing.T) {
	dec := i18n.NewJsonDecoder()
	data := make(map[string]any)
	if err := dec(strings.NewReader(`[{"x":"x123"}]`), &data); err != nil {
		t.Fatal(err)
	}
	if v, ok := data["x"]; ok {
		if v != "x123" {
			t.Fail()
		}
	} else {
		t.Fatal("JosnDecoder failed")
	}
}
