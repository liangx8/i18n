package i18n_test

import (
	"strings"
	"testing"

	"github.com/liangx8/i18n"
)

func Test_json(t *testing.T) {
	dec := i18n.NewJsonDecoder()
	data := make(map[string]any)
	dec(strings.NewReader(`{"x":"x123"}`), &data)
	t.Fatal(data)
}
