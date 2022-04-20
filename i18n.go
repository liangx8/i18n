package i18n

import (
	"fmt"
	"io"
	"strconv"
)

type (
	Decode func(io.Reader, *map[string]any) error
	locale struct {
		langs   []Lang
		storage map[string]any
	}
	Resource func() (io.Reader, error)

	Lang struct {
		l        string
		ldesc    string
		dec      Decode
		resource []Resource
	}
)

var localeStorage *locale

func reset() {
	localeStorage.langs = make([]Lang, 0)
	localeStorage.storage = make(map[string]any)
}
func Tr(key string, args ...any) string {
	if localeStorage == nil {
		return key
	}
	if val, ok := localeStorage.storage[key]; ok {
		switch vv := val.(type) {
		case int:
			return strconv.Itoa(vv)
		case string:
			if len(args) > 0 {
				return fmt.Sprintf(vv, args...)
			} else {
				return vv
			}
		default:
			panic("Tr(): doesn't expect reaching here")
		}
	}
	return key
}

// for resouce
func (lang *Lang) AddResource(res Resource) {
	lang.resource = append(lang.resource, res)
}
func (lang *Lang) Load(da *map[string]any) error {
	for _, r := range lang.resource {
		b, err := r()
		if err != nil {
			return err
		}
		if err := lang.dec(b, da); err != nil {
			return err
		}
	}
	return nil
}
