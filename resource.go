package i18n

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
)

func ResourceFilename(name string) Resource {
	return func() (io.Reader, error) {
		buf, err := ioutil.ReadFile(name)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(buf), nil
	}
}
func ResourceBytes(buffer []byte) Resource {
	return func() (io.Reader, error) {
		return bytes.NewBuffer(buffer), nil
	}
}
func ResourceText(text string) Resource {
	return func() (io.Reader, error) {
		return strings.NewReader(text), nil
	}
}
