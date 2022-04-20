package i18n

import (
	"encoding/json"
	"io"
	"log"
)

func NewJsonDecoder() Decode {
	return func(r io.Reader, da *map[string]any) error {
		dec := json.NewDecoder(r)
		t, err := dec.Token()
		if err != nil {
			return err
		}
		log.Printf("%T: %v", t, t)
		for dec.More() {
			if err := dec.Decode(da); err != nil {
				return err
			}

		}
		t, err = dec.Token()
		if err != nil {
			return err
		}
		log.Printf("%T: %v", t, t)
		return nil
	}
}
