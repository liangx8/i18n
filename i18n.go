package i18n

import (
	"fmt"
	"io"
	"log"
	"strconv"
)

type (
	Decode func(io.Reader, *map[string]any) error
	locale struct {
		langs   []*language
		storage map[string]any
	}
	Resource func() (io.Reader, error)

	language struct {
		l        string
		ldesc    string
		dec      Decode
		resource []Resource
	}
)

var localeStorage *locale

func init() {
	localeStorage = &locale{}
	localeStorage.langs = make([]*language, 0)
	localeStorage.storage = make(map[string]any)
}
func AddResource(lang string, res Resource) error {
	没找到 := true
	for _, l := range localeStorage.langs {
		if l.l == lang {
			l.AddResource(res)
			没找到 = false
		}
	}
	if 没找到 {
		return fmt.Errorf("Lang %s not found!", lang)
	}
	return nil
}
func Register(name, desc string, dec Decode) error {
	for _, l := range localeStorage.langs {
		if l.l == name {
			return fmt.Errorf("Lang %s is exists", name)
		}
	}
	localeStorage.langs = append(localeStorage.langs, &language{l: name, ldesc: desc, dec: dec})
	return nil

}
func Tr(key string, args ...any) string {
	if localeStorage == nil {
		panic("internal error at package i18n")
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
func SetLang(name string) error {
	for _, l := range localeStorage.langs {
		if l.l == name {
			l.Load(&localeStorage.storage)
			log.Printf("storage:%v", localeStorage.storage)
			return nil
		}
	}
	return fmt.Errorf("Lang %s not found!", name)
}

// for resouce
func (lang *language) AddResource(res Resource) {
	lang.resource = append(lang.resource, res)
}
func (lang *language) Load(da *map[string]any) error {
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
