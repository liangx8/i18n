package i18n

import (
	"encoding/json"
	"io"
)

func NewJsonDecoder() Decode {
	return func(r io.Reader, da *map[string]any) error {
		dec := json.NewDecoder(r)
		if err := dec.Decode(da); err != nil {
			return err
		}
		return nil
	}
}
