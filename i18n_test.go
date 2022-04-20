package i18n_test

import (
	"testing"

	"github.com/liangx8/i18n"
)

func TestRegister(t *testing.T) {
	if err := i18n.Register("zh_CN", "Simplified Chinese", i18n.NewJsonDecoder()); err != nil {
		t.Fatal(err)
	}
	if err := i18n.Register("zh_CN", "Simplified Chinese", i18n.NewJsonDecoder()); err == nil {
		t.Fatal("expected a error")
	}
	if err := i18n.AddResource("zh_CN", i18n.ResourceText("{'x':'交叉'}")); err != nil {
		t.Fatal(err)
	}
	if err := i18n.SetLang("zh_CN"); err != nil {
		t.Fatalf("SetLang():%v", err)
	}
}
func TestTr(t *testing.T) {
	TestRegister(t)
	t.Log(i18n.Tr("x"))
	t.Fail()
}
